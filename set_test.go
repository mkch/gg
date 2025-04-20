package gg_test

import (
	"testing"

	"github.com/mkch/gg"
)

func Test_Set(t *testing.T) {
	set := make(gg.Set[int])
	if set.Contains(1) {
		t.Fatalf("should not contain 1")
	}
	set.Add(1)
	if !set.Contains(1) {
		t.Fatalf("should contain 1")
	}
	set.Delete(1)
	if set.Contains(1) {
		t.Fatal("should not contain 1")
	}
}
