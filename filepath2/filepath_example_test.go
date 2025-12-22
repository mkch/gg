//go:build !windows

package filepath2_test

import (
	"fmt"

	"github.com/mkch/gg/filepath2"
)

func ExampleChangeExt() {
	fmt.Println(filepath2.ChangeExt("a.txt", "html"))
	fmt.Println(filepath2.ChangeExt("/path/file", ".c"))
	fmt.Println(filepath2.ChangeExt("a.txt", ""))
	fmt.Println(filepath2.ChangeExt(".a", "txt"))
	// Output:
	// a.html
	// /path/file.c
	// a
	// .txt
}
