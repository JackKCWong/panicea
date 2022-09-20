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

func Try(fn func()) (reErr error) {
	defer Catch(&reErr, nil)
	fn()

	return nil
}

func Trap[T any](fn func() T) (val T, reErr error) {
	defer Catch(&reErr, nil)

	return fn(), nil
}

func Catch(cause *error, handler func(error)) {
	if caught := recover(); caught != nil {
		if err, ok := caught.(error); ok {
			*cause = err
			if handler != nil {
				handler(err)
			}
		} else {
			panic(fmt.Errorf("expecting a error type but was: %q", caught))
		}
	}
}
