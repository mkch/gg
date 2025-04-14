package gg_test

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/mkch/gg"
)

func TestIf(t *testing.T) {
	if v := gg.If(1 == 1+0, "Yes", "No"); v != "Yes" {
		t.Fatal(v)
	}

	if v := gg.If(3%2 == 0, 1, 0); v != 0 {
		t.Fatal(v)
	}
}

func TestIfFunc(t *testing.T) {
	var (
		truePartExec, falsePareExec bool
	)

	if v := gg.IfFunc(
		true,
		func() string {
			truePartExec = true
			return "true"
		},
		func() string {
			falsePareExec = true
			return "false"
		}); v != "true" {
		t.Fatal(v)
	} else if !truePartExec {
		t.Fatal("true part should be executed")
	} else if falsePareExec {
		t.Fatal("false pare should not be executed")
	}
}

func ExampleMust() {
	fmt.Print(gg.Must(strconv.Atoi("1")))
	// Output:
	// 1
}

func ExampleMustOK() {
	gg.MustOK(os.Setenv("some_key_for_test", "some value"))
}

func TestMustOK(t *testing.T) {
	errInvalidArgument := errors.New("must >= 0")
	defer func() {
		if err := recover(); err != errInvalidArgument {
			t.Fatal(err)
		}
	}()
	greaterThanZero := func(i int) error {
		if i < 0 {
			return errInvalidArgument
		}
		return nil
	}

	gg.MustOK(greaterThanZero(-1))
}

func TestChainError(t *testing.T) {
	var err error
	var err1 = errors.New("err1")
	var err2 = errors.New("err2")
	gg.ChainError(func() error { return err1 }, &err)
	if !errors.Is(err, err1) {
		t.Fatalf("should contain err1")
	}
	if errors.Is(err, err2) {
		t.Fatalf("should not contain err2")
	}

	gg.ChainError(func() error { return nil }, &err)
	if !errors.Is(err, err1) {
		t.Fatalf("should contain err1")
	}
	if errors.Is(err, err2) {
		t.Fatalf("should not contain err2")
	}

	gg.ChainError(func() error { return err2 }, &err)
	if !errors.Is(err, err1) {
		t.Fatalf("should contain err1")
	}
	if !errors.Is(err, err2) {
		t.Fatalf("should contain err2")
	}
	t.Log(err)
}
