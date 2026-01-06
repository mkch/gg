package errorcheck

import (
	"errors"
	"testing"
)

func TestMustInvokesHandlerAndPanics(t *testing.T) {
	var handled bool
	handler := func(err error) {
		handled = true
	}

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Must should panic when err is non-nil")
		}
		if !handled {
			t.Fatalf("handler should be invoked before panic")
		}
	}()

	_ = Must(handler, 0, errors.New("boom"))
}
