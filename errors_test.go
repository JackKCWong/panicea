package panicea

import (
	"fmt"
	"testing"

)

func TestCheck(t *testing.T) {
	errorCaught := false
	err := func() (err error) {
		defer Catch(&err, func(err error) {
			errorCaught = true
		})

		Check(fmt.Errorf("I am an error"))

		return nil
	}()

	if err.Error() != "I am an error" {
		t.Logf("unexpected error: %+v", err)
		t.FailNow()
	}

	if !errorCaught {
		t.Log("error handler not called")
		t.FailNow()
	}

	t.Logf("%+v", err)
}
