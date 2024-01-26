// Package gg is a set of useful golang utilities.
package gg

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
