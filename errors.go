package panicea

import (
	"fmt"

	"github.com/pkg/errors"
)

func Must[T any](val T, err error) T {
	Check(err)

	return val
}

func Check(err error) {
	if err != nil {
		panic(errors.WithStack(err))
	}
}

func Catch(cause *error, handler func(error)) {
	if caught := recover(); caught != nil {
		if err, ok := caught.(error); ok {
			*cause = err 
			handler(err)
		} else {
			panic(fmt.Errorf("expecint an error but was: %q", caught))
		}
	}
}
