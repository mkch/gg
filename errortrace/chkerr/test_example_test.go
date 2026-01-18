package chkerr_test

import (
	"strconv"
	"testing"

	"github.com/mkch/gg/errortrace/chkerr"
)

func ExampleTest() {
	Test_Add := func(t *testing.T) {
		n := chkerr.Test(strconv.Atoi("1")).Must(t, "wrong number")
		sum := n + 10
		if sum != 11 {
			t.Fatal("sum is incorrect:", sum)
		}
	}
	_ = Test_Add
}
