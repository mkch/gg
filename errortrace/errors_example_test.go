package errortrace_test

import (
	"fmt"
	"io"
	"os"

	"github.com/mkch/gg/errortrace"
)

func openFile(name string) (io.WriteCloser, error) {
	f, err := os.OpenFile(name, os.O_WRONLY, 0222)
	if err != nil {
		return nil, errortrace.WithStack(err)
	}
	return f, err
}

func genFile(data string) error {
	w, err := openFile("no_such_file")
	if err != nil {
		return errortrace.ErrorfStack("can't write file: %w", err)
	}
	_, err = io.WriteString(w, data)
	if err != nil {
		return err
	}
	return w.Close()
}

func ExampleWithStack() {
	err := genFile("hello, world")
	if err != nil {
		err = errortrace.WithFileLine(err)
		fmt.Printf("%+v", err)
	}
	/* Output something like:

	   can't write file: open no_such_file: no such file or directory

	   github.com/mkch/gg/errortrace_test.ExampleWithStack()
	       path/errors_example_test.go:34

	       Caused by:
	       can't write file: open no_such_file: no such file or directory

	       ===== STACK TRACE =====
	       github.com/mkch/gg/errortrace_test.writeFile()
	           path/errors_example_test.go:22
	       github.com/mkch/gg/errortrace_test.ExampleWithStack()
	           path/errors_example_test.go:32
	       testing.runExample()
	           path/go/src/testing/run_example.go:63
	       ...
	       =======================

	           Caused by:
	           open no_such_file: no such file or directory

	           ===== STACK TRACE =====
	           github.com/mkch/gg/errortrace_test.openFile()
	               path/errors_example_test.go:14
	           github.com/mkch/gg/errortrace_test.writeFile()
	               path/errors_example_test.go:20
	           github.com/mkch/gg/errortrace_test.ExampleWithStack()
	               path/errors_example_test.go:32
	           testing.runExample()
	               path/go/src/testing/run_example.go:63
	           ...
	           =======================
	*/
}
