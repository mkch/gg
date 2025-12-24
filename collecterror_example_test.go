package gg_test

import (
	"errors"
	"fmt"
	"io"

	"github.com/mkch/gg"
)

// errWrite simulates an error returned from Write.
var errWrite = errors.New("write failed")

// errClose simulates an error returned from Close.
var errClose = errors.New("close failed")

// mockFile simulates a WriteCloser that always returns errors on Write and Close.
type mockFile struct{}

func Open() (io.WriteCloser, error) {
	return &mockFile{}, nil
}

func (m mockFile) Write(p []byte) (int, error) {
	return 0, errWrite
}

func (m mockFile) Close() error {
	return errClose
}

// FileOp opens a file, writes to it, and ensures proper error chaining.
func FileOp() (err error) {
	f, err := Open()
	if err != nil {
		return err
	}
	// Ensure f is closed and any close error is combined with err.
	defer gg.CollectError(f.Close, &err)

	_, err = f.Write([]byte("x"))
	return err
}

func ExampleCollectError() {
	err := FileOp()
	// err contains both write and close errors.
	fmt.Println(err)
	fmt.Println(errors.Is(err, errWrite), errors.Is(err, errClose))

	// Output:
	// write failed
	// close failed
	// true true
}
