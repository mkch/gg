// Package filepath implements utility routines for manipulating filename paths.
package filepath2

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestChangeExt(t *testing.T) {
	type args struct {
		path   string
		newExt string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal", args{"a.txt", "html"}, "a.html"},
		{"no ext", args{"abc", ".c"}, "abc.c"},
		{"empty new ext", args{"abc", ""}, "abc"},
		{"multi dots", args{"/path/a.b/c.e", "f"}, filepath.FromSlash("/path/a.b/c.f")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeExt(tt.args.path, tt.args.newExt); got != tt.want {
				t.Errorf("ChangeExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"normal", "a.txt", "a"},
		{"path", "/path/a.txt", "a"},
		{"multi dots", "/path/a.b.txt", "a.b"},
		{"empty name", ".x", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Name(tt.arg); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleName() {
	fmt.Println(Name("a.txt"))
	fmt.Println(Name("/path/b.c.e"))
	fmt.Printf("%#v", Name(".c"))
	// Output:
	// a
	// b.c
	// ""
}
