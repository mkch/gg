package pkg

import "github.com/mkch/gg"

type myStruct struct {
	_      gg.NoCopy
	Field1 int
	Field2 string
}

func f() {
	var a myStruct
	_ = a
}
