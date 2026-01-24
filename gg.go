// Package gg is a set of useful golang utilities.
package gg

import (
	"errors"
	"unsafe"
)

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

// Zero returns the zero value of type T.
func Zero[T any]() (zero T) {
	return
}

// Clear sets the value pointed to by ptr to its zero value.
func Clear[T any](ptr *T) {
	// Converts ptr to a pointer to a one-element array,
	// then slices it to get a []T of length 1,
	// and finally calls clear on the slice.
	clear((*[1]T)(unsafe.Pointer(ptr))[:])
}

// clearSafe is a safe version of Clear without using unsafe.
//
// The benchmark shows that clearSafe is almost as fast as Clear
// for small sized types(< 10240 bytes), but significantly
// slower for large sized types.
// The benchmark also shows that clearSafe do heap allocation for
// large sized types.
//
// The benchmark results:
//
//	go test -run ^BenchmarkClear -bench . -benchmem
//	goos: darwin
//	goarch: amd64
//	pkg: github.com/mkch/gg
//	cpu: Intel(R) Core(TM) i7-7920HQ CPU @ 3.10GHz
//	BenchmarkClearSafeSmall-8         718131              1499 ns/op               0 B/op          0 allocs/op
//	BenchmarkClearSmall-8             752520              1410 ns/op               0 B/op          0 allocs/op
//	BenchmarkClearSafeLarge-8          10695            111626 ns/op          819215 B/op          1 allocs/op
//	BenchmarkClearLarge-8              75961             15195 ns/op               0 B/op          0 allocs/op
//	PASS
//	ok      github.com/mkch/gg      4.961s
func clearSafe[T any](ptr *T) {
	var zero T
	*ptr = zero
}
