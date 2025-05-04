package os2_test

import (
	"bytes"
	"testing"

	"os"

	"github.com/mkch/gg/os2"
)

func TestCopyFil(t *testing.T) {
	const src = "testdata/a.txt"
	const dest = "testdata/b.txt"
	defer os.Remove(dest)
	err := os2.CopyFile(src, dest, true)
	if err != nil {
		t.Fatal(err)
	}
	srcBytes, err := os.ReadFile(src)
	if err != nil {
		t.Fatal(err)
	}
	destBytes, err := os.ReadFile(dest)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(srcBytes, destBytes) {
		t.Fatal(string(destBytes))
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		t.Fatal(err)
	}
	destInfo, err := os.Stat(dest)
	if err != nil {
		t.Fatal(err)
	}
	if destMode := destInfo.Mode(); destMode != srcInfo.Mode() {
		t.Fatalf("0%o", destMode)
	}
}
