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

// CollectError executes f and collects its error into *dest.
// It is intended for use with defer so cleanup errors are not lost.
// If f returns an error and *dest is nil, store it in *dest;
// otherwise join it with *dest using [errors.Join].
// Argument dest can't be nil, and it usually points to the function's named return error.
func CollectError(f func() error, dest *error) {
	err := f()
	if err == nil {
		return
	}
	*dest = IfFunc(*dest == nil,
		func() error { return err },
		func() error { return errors.Join(*dest, err) },
	)
}
