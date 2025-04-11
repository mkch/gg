package sgo

import (
	"context"
	"slices"
	"sync/atomic"
	"testing"
)

func TestGroup(t *testing.T) {
	var result [2]int
	NewGroup().
		Go(func() {
			result[0] = 1
		}).
		Go(func() {
			result[1] = 2
		}).
		Wait()

	slices.Sort(result[:])
	if !slices.Equal(result[:], []int{1, 2}) {
		t.Fatal(result)
	}
}

func TestGroupMaxGo(t *testing.T) {
	group := NewGroup().MaxGo(3)
	var count atomic.Int32
	group.Go(func() {
		count.Add(1)
		defer count.Add(-1)
		if n := count.Load(); n > 3 {
			panic(n)
		}
	}).Go(func() {
		count.Add(1)
		defer count.Add(-1)
		if n := count.Load(); n > 3 {
			panic(n)
		}
	}).Go(func() {
		count.Add(1)
		defer count.Add(-1)
		if n := count.Load(); n > 3 {
			panic(n)
		}
	}).Go(func() {
		count.Add(1)
		defer count.Add(-1)
		if n := count.Load(); n > 3 {
			panic(n)
		}
	}).Go(func() {
		count.Add(1)
		defer count.Add(-1)
		if n := count.Load(); n > 3 {
			panic(n)
		}
	})
	group.Wait()
}

func TestRacer(t *testing.T) {
	type result struct {
		Err error
		N   int
	}
	racer := NewRacer[result](context.Background())

	racer.Go(func(ctx context.Context) result {
		select {
		case <-ctx.Done():
			return result{context.Cause(ctx), 0}
		default:
			return result{nil, 5050}
		}
	}).Go(func(ctx context.Context) result {
		var sum int
		for i := range 100 {
			sum += i
			select {
			case <-ctx.Done():
				return result{context.Cause(ctx), 0}
			default:
			}
		}
		return result{nil, sum}
	})

	r := racer.Proclaim()
	if r.Err != nil || r.N != 5050 {
		t.Fatal(r)
	}
}

func TestCollect(t *testing.T) {
	result := NewCollector[int]().
		Go(func() int { return 1 }).
		Go(func() int { return 2 }).
		Collect()
	slices.Sort(result)
	if !slices.Equal(result, []int{1, 2}) {
		t.Fatal(result)
	}
}

func TestEmptyWait(t *testing.T) {
	NewGroup().Wait()
}

func TestEmptyCollect(t *testing.T) {
	result := NewCollector[struct{}]().Collect()
	if len(result) != 0 {
		t.Fatal(result)
	}
}

func TestEmptyRace(t *testing.T) {
	result := NewRacer[int](context.Background()).Proclaim()
	if result != 0 {
		t.Fatal(result)
	}
}
