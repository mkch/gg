package errortrace

import (
	"errors"
	"fmt"
	"path"
	"regexp"
	"testing"

	"github.com/mkch/gg/runtime2"
)

func Test_Errorf(t *testing.T) {
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
	err1 := ErrorfStack("Error caused by %w and %w", errors.New("first cause"), errors.New("second cause"))

	f := err1.StackFrames()
	if len(f.Frames) < 6 {
		t.Fatal(f)
	}
	for i := range 5 {
		// ".f1", ".f2" etc.
		if funcName := path.Ext(f.Frames[i].Function); funcName != fmt.Sprintf(".f%v", i+1) {
			t.Fatal(f)
		}
	}

	const expected = `Error caused by first cause and second cause

===== STACK TRACE =====
github.com/mkch/gg/errortrace.f1()
	<this_file>:34
github.com/mkch/gg/errortrace.f2()
	<this_file>:30
github.com/mkch/gg/errortrace.f3()
	<this_file>:26
github.com/mkch/gg/errortrace.f4()
	<this_file>:22
github.com/mkch/gg/errortrace.f5()
	<this_file>:18
github.com/mkch/gg/errortrace.Test_Errorf()
	<this_file>:14
=======================
`

	output := fmt.Sprintf("%+v", err1)
	// Replace file path with <this_file> for testing.
	thisFile := runtime2.Source().File // Full path of this file.
	re := regexp.MustCompile(`(\t*.+?\(\)\n\t*)(.+)\:(\d+\n)`)
	output = re.ReplaceAllStringFunc(output, func(str string) string {
		m := re.FindStringSubmatch(str)
		// Remove all
		//
		// pkg.f()
		// 	path_to_file:line
		//
		// which path_to_file is not this file.
		if m[2] != thisFile {
			return ""
		}
		// Replace all occurrences of this file path with <this_file>.
		return fmt.Sprintf("%s<this_file>:%s", m[1], m[3])
	})
	if output != expected {
		t.Fatalf("output did not match expected:\n%s", output)
	}
}
