package gg_test

import (
	"os"

	"github.com/mkch/gg"
)

func WriteFile(name string, data []byte) (err error) {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer gg.ChainError(f.Close, &err)
	_, err = f.Write(data)
	return err
}

func ExampleChainError() {
	WriteFile("file", []byte("data"))
}
