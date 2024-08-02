package gg_test

import (
	"path/filepath"
	"testing"

	"github.com/mkch/gg"
)

func TestSource(t *testing.T) {
	src := gg.Source()
	// Modify the lie number below if changed.
	if filepath.Base(src.File) != "source_test.go" || src.Line != 11 || src.Function != "github.com/mkch/gg_test.TestSource" {
		t.Fatal(src)
	}
}
