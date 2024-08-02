package runtime_test

import (
	"path"
	"path/filepath"
	"testing"

	"github.com/mkch/gg/runtime"
)

func TestSource(t *testing.T) {
	src := runtime.Source()
	// Modify the lie number below if changed.
	if filepath.Base(src.File) != "source_test.go" || src.Line != 12 || path.Base(src.Function) != "runtime_test.TestSource" {
		t.Fatal(src)
	}
}
