package gg_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/mkch/gg"
)

func TestIf(t *testing.T) {
	if v := gg.If(1 == 1+0, "Yes", "No"); v != "Yes" {
		t.Fatal(v)
	}

	if v := gg.If(3%2 == 0, 1, 0); v != 0 {
		t.Fatal(v)
	}
}

func ExampleMust() {
	fmt.Print(gg.Must(strconv.Atoi("1")))
	// Output:
	// 1
}
