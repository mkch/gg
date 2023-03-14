// Package gg is a set of useful golang utilities.
package gg

// If returns truePart of cond is true, or returns falsePart.
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

// Must returns v is err is nil, or it panic with err.
// It is useful to wrap some functions with 2 return values
// when there is no better way to handle errors than panicking.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
