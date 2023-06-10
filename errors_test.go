package panicea

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCheck(t *testing.T) {
	errorCaught := false
	err := func() (err error) {
		defer catch(&err, func(err error) {
			errorCaught = true
		})

		panic(fmt.Errorf("I am an error"))
	}()

	if err.Error() != "I am an error" {
		t.Fatalf("unexpected error: %+v", err)
	}

	if !errorCaught {
		t.Fatal("error handler not called")
	}

	t.Logf("%+v", err)
}

func TestTry(t *testing.T) {
	n, err := Try(func() int {
		return 10
	})

	if err != nil {
		t.Fatalf("unexpected error: %q", err)
	}

	if n != 10 {
		t.Fatalf("unexpected result: %q", n)
	}

	n, err = Try(func() int {
		panic(fmt.Errorf("I am an error"))
	})

	if err == nil || err.Error() != "I am an error" {
		t.Fatalf("unexpected error: %q", err)
	}

	if n != 0 {
		t.Fatalf("unexpected result: %q", n)
	}

	t.Logf("%+v", err)
}

func TestCatch(t *testing.T) {
	f1 := func() (int, error) {
		return 10, nil
	}

	v := Catch(f1()).Throw("failure!")

	if v != 10 {
		t.Fatalf("expected val: %v", v)
	}

	f2 := func() (int, error) {
		return 0, fmt.Errorf("random error")
	}

	err := Trap(func() {
		Catch(f2()).Throw("failure: %w")
	})

	if err.Error() != "failure: random error" {
		t.Fatalf("unexpected err: %v", err)
	}

	err = Trap(func() {
		Catch(0, fmt.Errorf("boom!")).Throw(BadRequest, "bad input: %w")
	})

	if httpErr, ok := err.(*HttpError); !ok {
		t.Fatalf("unexpected err: %v", err)
	} else {
		if httpErr.StatusCode != http.StatusBadRequest {
			t.Fatalf("unexpected err: %v", err)
		}

		if httpErr.Error() != "status=400, bad input: boom!" {
			t.Fatalf("unexpected err: %v", err)
		}
	}
}
