package errors2_test

import (
	"fmt"
	"io"
	"os"

	"github.com/mkch/gg/errors2"
)

func openFile(name string) (io.WriteCloser, error) {
	f, err := os.OpenFile(name, os.O_WRONLY, 0222)
	if err != nil {
		return nil, errors2.WithStack(err)
	}
	return f, err
}

func writeFile(f io.WriteCloser, data string) error {
	w, err := openFile("no_such_file")
	if err != nil {
		return err
	}
	_, err = io.WriteString(w, data)
	if err != nil {
		return err
	}
	return w.Close()
}

func ExampleWithStack() {
	err := writeFile(nil, "hello, world")
	if err != nil {
		fmt.Printf("%+v", err)
	}
	// Will output something like:
	//
	//*errors2.ErrorWithStack: open no_such_file: The system cannot find the file specified.
	//
	//===== STACK TRACE =====
	//github.com/mkch/gg/errors2_test.openFile()
	//	path/to/errors_example_test.go:14
	//github.com/mkch/gg/errors2_test.writeFile()
	//	path/to/errors_example_test.go:20
	//github.com/mkch/gg/errors2_test.ExampleWithStack()
	//	path/to/errors_example_test.go:32
	//testing.runExample()
	//	path/to/Go/src/testing/run_example.go:63
	//testing.runExamples()
	//	path/to/Go/src/testing/example.go:41
	//testing.(*M).Run()
	//	path/to/Go/src/testing/testing.go:2339
	// main.main()
	//	testmain.go:51
	// runtime.main()
	//	path/to/Go/src/runtime/proc.go:285
	// runtime.goexit()
	//	path/to/Go/src/runtime/asm_amd64.s:1693
	//=======================

	//Caused by:
	//	*fs.PathError: open no_such_file: The system cannot find the file specified.

	//Caused by:
	//	*syscall.Errno: The system cannot find the file specified.
}
