// Package sorted provides functions for manipulating sorted slices.
package sorted

import (
	"cmp"
	"slices"
)

// BisectRight searches for the insertion point for x in a sorted slice.
// The return value is the index where x would be inserted
// such that all elements in s[:i] are <= x, and all elements in s[i:] are > x.
func BisectRight[S ~[]E, E cmp.Ordered](s S, target E) int {
	// BinarySearchFunc searches for the smallest index i where cmp(s[i], x) >= 0.
	// We want the first index where s[i] > x.
	// Our comparison returns 1 if elem > target (>=0 for BinarySearchFunc),
	// and -1 if elem <= target (<0 for BinarySearchFunc).
	i, _ := slices.BinarySearchFunc(s, target, func(elem, target E) int {
		if elem > target {
			return 1
		}
		return -1
	})

	return i
}

// BisectRightFunc works like [BisectRight], but uses a custom comparison function.
func BisectRightFunc[S ~[]E, E, T any](s S, target T, cmp func(E, T) int) int {
	i, _ := slices.BinarySearchFunc(s, target, func(e E, t T) int {
		if cmp(e, t) > 0 {
			return 1
		}
		return -1
	})
	return i
}

// Insert inserts an element e into a ascendingly sorted slice s,
// such that the elements in the updated slice remain in ascending order.
// The insertion position is determined by [slices.BinarySearch]: inserting before any existing elements equal to e.
func Insert[S ~[]E, E cmp.Ordered](s S, e E) S {
	i, _ := slices.BinarySearch(s, e)
	return slices.Insert(s, i, e)
}

// InsertFunc works like [Insert], but use a custom comparison function.
// The slice must be sorted in increasing order, where "increasing" is defined
// the same way as in [slices.BinarySearchFunc].
func InsertFunc[E any, S ~[]E](s S, e E, cmp func(a, b E) int) S {
	i, _ := slices.BinarySearchFunc(s, e, cmp)
	return slices.Insert(s, i, e)
}

// Delete deletes all elements e from a ascendingly sorted slice s.
func Delete[S ~[]E, E cmp.Ordered](s S, e E) S {
	i, exists := slices.BinarySearch(s, e)
	if !exists {
		return s
	}
	j := BisectRight(s, e)
	return slices.Delete(s, i, j)
}

// DeleteFunc works like [Delete], but use a custom comparison function.
// The slice must be sorted in increasing order, where "increasing" is defined
// the same way as in [slices.BinarySearchFunc].
func DeleteFunc[E any, S ~[]E](s S, e E, cmp func(a, b E) int) S {
	i, exists := slices.BinarySearchFunc(s, e, cmp)
	if !exists {
		return s
	}
	j := BisectRightFunc(s, e, cmp)
	return slices.Delete(s, i, j)
}

// Find returns a run of elements equal to e from an ascendingly sorted slice s.
// The content of returned slice must not be modified; it is valid only until the next update of s
func Find[S ~[]E, E cmp.Ordered](s S, e E) S {
	i, exists := slices.BinarySearch(s, e)
	if !exists {
		return nil
	}
	j := BisectRight(s, e)
	return s[i:j]
}

// FindFunc works like [Find], but use a custom comparison function.
// The slice must be sorted in increasing order, where "increasing" is defined
// the same way as in [slices.BinarySearchFunc].
func FindFunc[E any, S ~[]E](s S, e E, cmp func(a, b E) int) S {
	i, exists := slices.BinarySearchFunc(s, e, cmp)
	if !exists {
		return nil
	}
	j := BisectRightFunc(s, e, cmp)
	return s[i:j]
}

// Prefix returns the prefix run of ascendingly sorted slice s.
// The prefix run is the maximal sequence of identical elements from the start of s.
// If the length of s is 0, the returned prefix is nil.
func Prefix[E cmp.Ordered, S ~[]E](s S) S {
	if len(s) == 0 {
		return nil
	}
	j := BisectRight(s, s[0])
	return s[:j]
}

// PrefixFunc works like [Prefix], but use a custom comparison function.
// The slice must be sorted in increasing order, where "increasing" is defined
// the same way as in [slices.BinarySearchFunc].
func PrefixFunc[E any, S ~[]E](s S, cmp func(a, b E) int) S {
	if len(s) == 0 {
		return nil
	}
	j := BisectRightFunc(s, s[0], cmp)
	return s[:j]
}
