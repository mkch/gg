package slices

import (
	"strconv"
	"testing"
)

type Slice[E any] []E

func (s Slice[E]) Len() int {
	return len(s)
}

func TestFilter(t *testing.T) {
	src := Slice[int]{1, 2, 3, 4, 5}
	dest := Filter(src, func(v int) bool { return v%2 == 0 })
	if dest.Len() != 2 || dest[0] != 2 || dest[1] != 4 {
		t.Fatal(dest)
	}
}

func TestMap(t *testing.T) {
	src := Slice[int]{1, 2, 3}
	var dest = Map(src, func(n int) int { return n + 1 })
	if len(dest) != 3 || dest[0] != 2 || dest[1] != 3 || dest[2] != 4 {
		t.Fatal(dest)
	}
}

func TestReduce(t *testing.T) {
	r := Reduce([]int{1, 2, 3, 4, 5},
		func(a string, v int, i int) string { return a + strconv.Itoa(v) },
		"str:")
	if r != "str:12345" {
		t.Fatal(r)
	}
}
