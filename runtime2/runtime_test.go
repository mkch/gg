package runtime2_test

import (
	"fmt"
	"path"
	"testing"

	"github.com/mkch/gg/runtime2"
)

func TestStack(t *testing.T) {
	f5(t)
}

func f5(t *testing.T) {
	f4(t)
}

func f4(t *testing.T) {
	f3(t)
}

func f3(t *testing.T) {
	f2(t)
}

func f2(t *testing.T) {
	f1(t)
}

func f1(t *testing.T) {
	f := runtime2.Stack(0, 0)
	if len(f.Frames) < 6 {
		t.Fatal(f)
	}
	for i := range 5 {
		// ".f1", ".f2" etc.
		if funcName := path.Ext(f.Frames[i].Function); funcName != fmt.Sprintf(".f%v", i+1) {
			t.Fatal(f)
		}
	}
}
