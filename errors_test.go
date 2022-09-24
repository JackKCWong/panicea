package panicea

import (
	"fmt"
	"testing"
)

func TestCheck(t *testing.T) {
	errorCaught := false
	err := func() (err error) {
		defer catch(&err, func(err error) {
			errorCaught = true
		})

		Check(fmt.Errorf("I am an error"))

		return nil
	}()

	if err.Error() != "I am an error" {
		t.Fatalf("unexpected error: %+v", err)
	}

	if !errorCaught {
		t.Fatal("error handler not called")
	}

	t.Logf("%+v", err)
}

func TestTrap(t *testing.T) {
	n, err := Trap(func() int {
		return 10
	})

	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if n != 10 {
		t.Fatalf("unexpected result: %q", n)
	}

	n, err = Trap(func() int {
		Check(fmt.Errorf("I am an error"))

		return -1
	})

	if err == nil || err.Error() != "I am an error" {
		t.Fatalf("unexpected error: %q", err)
	}

	if n != 0 {
		t.Fatalf("unexpected result: %q", n)
	}

	t.Logf("%+v", err)
}
