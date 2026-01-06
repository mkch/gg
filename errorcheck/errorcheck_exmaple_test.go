package errorcheck_test

import (
	"fmt"
	"net"
	"strconv"

	"github.com/mkch/gg/errorcheck"
)

// errorHandler is a sample error handler that prints the error.
func errorHandler(err error) {
	fmt.Printf("Error occurred: %v\n", err)
}

// Must is a wrapper around [errorcheck.Must] using [errorHandler].
func Must[T any](v T, err error) T {
	return errorcheck.Must(errorHandler, v, err)
}

// Must is a wrapper around [errorcheck.Must2] using [errorHandler].
func Must2[T1 any, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	return errorcheck.Must2(errorHandler, v1, v2, err)
}

func ExampleMust() {
	// Must eliminates explicit error checking, e.g.:
	//   val, err := strconv.Atoi("123")
	//   if err != nil { ... }
	//   fmt.Println(val + 333)
	// becomes:
	fmt.Println(Must(strconv.Atoi("123")) + 333)

	// Must2 works similarly for functions returning two values plus error
	fmt.Println(Must2(net.SplitHostPort("localhost:8080")))

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	// When error occurs, handler is called before panic
	fmt.Println(Must(strconv.Atoi("abc")))
	// Output:
	// 456
	// localhost 8080
	// Error occurred: strconv.Atoi: parsing "abc": invalid syntax
	// panic: strconv.Atoi: parsing "abc": invalid syntax
}
