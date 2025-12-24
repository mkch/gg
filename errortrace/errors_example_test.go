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
		// Add stack trace.
		return nil, errortrace.WithStack(err)
	}
	return f, err
}

func genFile(data string) error {
	w, err := openFile("no_such_file")
	if err != nil {
		// Prepend context and add stack trace.
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
		// Wrapped Error can still be printed by errortrace.Fprint.
		err = fmt.Errorf("%w", err)
		errortrace.Fprint(os.Stdout, err)
	}
	/* Output something like:

	can't write file: open no_such_file: no such file or directory

	github.com/mkch/gg/errortrace_test.ExampleWithStack()
	    path/errors_example_test.go:36

	    Caused by:
	    can't write file: open no_such_file: no such file or directory

	    ===== STACK TRACE =====
	    github.com/mkch/gg/errortrace_test.writeFile()
	        path/errors_example_test.go:24
	    github.com/mkch/gg/errortrace_test.ExampleWithStack()
	        path/errors_example_test.go:34
	    testing.runExample()
	        path/go/src/testing/run_example.go:63
	    ...
	    =======================

	        Caused by:
	        open no_such_file: no such file or directory

	        ===== STACK TRACE =====
	        github.com/mkch/gg/errortrace_test.openFile()
	            path/errors_example_test.go:15
	        github.com/mkch/gg/errortrace_test.writeFile()
	            path/errors_example_test.go:21
	        github.com/mkch/gg/errortrace_test.ExampleWithStack()
	            path/errors_example_test.go:34
	        testing.runExample()
	            path/go/src/testing/run_example.go:63
	        ...
	        =======================
	*/
}
