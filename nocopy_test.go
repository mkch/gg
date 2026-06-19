package gg_test

import (
	"os/exec"
	"testing"
)

func TestNoCopy(t *testing.T) {
	cmd := exec.Command("go", "vet", "-copylocks", "testdata/testnocopy.go")
	err := cmd.Run()
	if err == nil {
		t.Fatal("expected go vet to fail, but it succeeded")
	}
}
