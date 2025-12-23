// Package gg is a set of useful golang utilities.
package gg

import "errors"

// If returns truePart if cond is true, or returns falsePart.
// Note that Go evaluates truePart and falsePart regardless of cond when calling If.
func If[T any](cond bool, truePart T, falsePart T) T {
	if cond {
		return truePart
	} else {
		return falsePart
	}
}

// IfFunc returns the result of truePart() if cond is true, or the result of falsePart().
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

// ChainError executes f and, if it returns a non-nil error, merges it into dest using errors.Join.
// It is intended for use with defer so cleanup errors are combined with the function's return error.
// If the return value of f is nil nothing changes; if *dest is nil it becomes the return value of f;
// otherwise *dest is set to errors.Join(*dest, return_of_f).
// Argument dest should point to the function's named return error.
func ChainError(f func() error, dest *error) {
	err := f()
	if err == nil {
		return
	}
	*dest = IfFunc(*dest == nil,
		func() error { return err },
		func() error { return errors.Join(*dest, err) },
	)
}
