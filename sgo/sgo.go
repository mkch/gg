// Package sgo implements structured concurrency in go.
package sgo

import (
	"context"
	"errors"
	"sync"
)

// Group is a set of concurrent tasks.
type Group struct {
	wg    sync.WaitGroup
	limit chan struct{}
}

// NewGroup creates a Group.
func NewGroup() *Group {
	return &Group{}
}

// MaxGo sets the maximum allowed concurrency
// (the number of goroutines running simultaneously) in this Group.
// If n<=0, the concurrency limit is unbounded.
// Returns the group itself to allow method chaining.
//
// The calls to MaxGo must happen before [Group.Go].
// If a group is reused to wait for several independent sets of tasks,
// new MaxGo calls must happen after all previous Wait calls have returned,
// and before any [Group.Go] calls.
func (group *Group) MaxGo(n int) *Group {
	if n <= 0 {
		group.limit = nil
	} else {
		group.limit = make(chan struct{}, n)
	}
	return group
}

// Go creates a concurrent task in the group.
// If the group has reached maximum concurrency, it waits until some task
// completes before running this one.
// Returns the group itself to allow method chaining.
//
// The calls to Go should execute before [Group.Wait].
// If a group is reused to wait for several independent sets of tasks,
// new Go calls must happen after all previous Wait calls have returned.
func (group *Group) Go(task func()) *Group {
	if group.limit != nil {
		group.limit <- struct{}{}
	}
	group.wg.Add(1)
	go func() {
		defer func() {
			if group.limit != nil {
				<-group.limit
			}
			group.wg.Done()
		}()
		task()
	}()
	return group
}

// Wait blocks until all tasks in the group complete.
func (group *Group) Wait() {
	group.wg.Wait()
}

// Collector runs a set of concurrent tasks and collects their execution results.
// The generic parameter T specifies the return type of the tasks.
type Collector[T any] struct {
	g       *Group
	results chan T
}

// NewCollector creates a Collector.
func NewCollector[T any]() *Collector[T] {
	return &Collector[T]{
		g:       NewGroup(),
		results: make(chan T),
	}
}

// MaxGo sets the maximum allowed concurrency
// See [Group.MaxGo].
func (collector *Collector[T]) MaxGo(n int) *Collector[T] {
	collector.g.MaxGo(n)
	return collector
}

// Go creates a concurrent task in the collector.
// If the collector has reached maximum concurrency, it waits until some task
// completes before running this one.
// The return value of the tasks will be returned by [Collector.Collect].
// For the calling sequence between this method and [Collector.Collect], see [Group.Go].
func (collector *Collector[T]) Go(task func() T) *Collector[T] {
	collector.g.Go(func() {
		collector.results <- task()
	})
	return collector
}

// Collect blocks until all tasks in the collector complete and returns their results.
// Note: The order of elements in results is nondeterministic.
//
// Empty slice will be returned if there is no [Collector.Go] call before this method.
func (collector *Collector[T]) Collect() (result []T) {
	join := make(chan struct{})
	// c.results has multiple writers, no one should close it.
	// The flowing goroutine makes join a signal that indicates
	// all writing finished.
	go func() {
		collector.g.Wait()
		close(join)
	}()

	for {
		select {
		case r := <-collector.results:
			result = append(result, r)
		case <-join:
			return
		}
	}
}

// ErrRaceFailed is the cause used when canceling context in [Racer].
var ErrRaceFailed = errors.New("race failed")

// Racer runs a set of concurrent tasks and makes them race(compete),
// where the first task to return a result wins and cancels the context.
type Racer[T any] struct {
	// Cleanup is used to clean up return values from failed tasks.
	// Can be nil, which means no cleanup is needed.
	// See [Racer.Collect].
	// Should be set before any [Racer.Go] calls.
	Cleanup func(T)
	g       *Group
	results chan T
	ctx     context.Context
	cancel  context.CancelCauseFunc
}

// MaxGo sets the maximum allowed concurrency.
// See [Group.MaxGo].
func (racer *Racer[T]) MaxGo(n int) *Racer[T] {
	racer.g.MaxGo(n)
	return racer
}

// NewRacer creates a [Racer].
// The ctx parameter will be passed to tasks when they run.
func NewRacer[T any](ctx context.Context) *Racer[T] {
	var ret = Racer[T]{
		g:       NewGroup(),
		results: make(chan T),
	}
	ret.ctx, ret.cancel = context.WithCancelCause(ctx)
	return &ret
}

// Go creates a concurrent task in the Racer.
func (racer *Racer[T]) Go(task func(context.Context) T) *Racer[T] {
	racer.g.Go(func() {
		racer.results <- task(racer.ctx)
	})
	return racer
}

// Proclaim blocks until a task in the racer returns (wins),
// at which point it cancels the context with [ErrRaceFailed] as the cause,
// and then returns the result of the winner.
// After a task wins, subsequent returns from other tasks
// (which may occur due to a tie or not respecting context cancellation)
// are cleaned up using racer.Cleanup if it is not nil.
//
// The zero value of T will be returned if there is no [Racer.Go] call before this method.
func (racer *Racer[T]) Proclaim() (result T) {
	// See [Collector.Collect]
	join := make(chan struct{})
	go func() {
		racer.g.Wait()
		close(join)
	}()

	select {
	case result = <-racer.results: // the winner is claimed.
		// cancels the context.
		racer.cancel(ErrRaceFailed)
		// cleanup the failures
		for {
			select {
			case failure := <-racer.results:
				if racer.Cleanup != nil {
					racer.Cleanup(failure)
				}
			// join tasks
			case <-join:
				return
			}
		}
	case <-join:
		return
	}
}

// Wait is equivalent to creating a [Group] with [NewGroup],
// executing all tasks via [Group.Go], and then calling [Group.Wait].
func Wait(maxGo int, tasks ...func()) {
	group := NewGroup().MaxGo(maxGo)
	for _, task := range tasks {
		group.Go(task)
	}
	group.Wait()
}

// Collect is equivalent to creating a [Collector] with [NewCollector],
// executing all tasks via [Collector.Go], and returning the result of [Collector.Collect].
func Collect[T any](maxGo int, tasks ...func() T) []T {
	collector := NewCollector[T]().MaxGo(maxGo)
	for _, task := range tasks {
		collector.Go(task)
	}
	return collector.Collect()
}

// Race is equivalent to creating a [Racer] with [NewRacer], setting its Cleanup,
// running all tasks via [Racer.Go], and returning the result of [Racer.Proclaim].
func Race[T any](ctx context.Context, maxGo int, cleanup func(T), tasks ...func(context.Context) T) T {
	racer := NewRacer[T](ctx).MaxGo(maxGo)
	racer.Cleanup = cleanup
	for _, task := range tasks {
		racer.Go(task)
	}
	return racer.Proclaim()
}
