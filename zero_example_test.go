package gg_test

import "github.com/mkch/gg"

type List[T any] struct {
	Value T
	Next  *List[T]
}

func FindFirstFunc[T any](list *List[T], f func(T) bool) (T, bool) {
	for l := list; l != nil; l = l.Next {
		if f(l.Value) {
			return l.Value, true
		}
	}
	// Of course, we can also use named return value.
	return gg.Zero[T](), false
}
