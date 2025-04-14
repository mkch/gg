// Package gg is a set of useful golang utilities.
package gg

import "fmt"

// If returns truePart if cond is true, or returns falsePart.
// Something like ternary operator in C:
//
//	cond ? truePart : falsePart
func If[T any](cond bool, truePart T, falsePart T) T {
	if cond {
		return truePart
	} else {
		return falsePart
	}
}

// IfFunc calls truePart() and returns it's return value if cond is true,
// or calls falsePart() and returns it's return value.
// If cond is true, falsePart will not be executed.
// if cond if false, truePart will not be executed.
// Something like ternary operator in C:
//
//	cond ? truePart() : falsePart()
func IfFunc[T any](cond bool, truePart func() T, falsePart func() T) T {
	if cond {
		return truePart()
	} else {
		return falsePart()
	}
}

// Must returns v if err is nil, or it panic with err.
// Must is useful to wrap a function call returning value and error
// when there is no better way to handle the error other than panicking.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// MustOK panics if err is not nil.
func MustOK(err error) {
	if err != nil {
		panic(err)
	}
}

// ChainError first executes f, and returns if it returns nil.
// When f's return value is not nil (counted as err), if *dest is nil, sets *dest to err,
// otherwise wraps *dest and err into one error and assigns it to *dest.
// ChainError can be used after the defer keyword, for closing [io.Closer] while preserving the error returned by [io.Closer.Close].
func ChainError(f func() error, dest *error) {
	ferr := f()
	if ferr == nil {
		return
	}
	*dest = IfFunc(*dest == nil,
		func() error { return ferr },
		func() error { return fmt.Errorf("%w; %w", *dest, ferr) },
	)
}
