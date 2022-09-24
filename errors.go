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
	defer catch(&reErr, nil)
	fn()

	return nil
}

func Trap[T any](fn func() T) (val T, reErr error) {
	defer catch(&reErr, nil)

	return fn(), nil
}

type Result[T any] struct {
	Val T	
	Err error
}

func Pack[T any] (val T, err error) Result[T] {
	return Result[T]{
		Val: val,
		Err: err,
	}
}

func Wrap[T any](fn func() T) (Result[T]) {
	var r Result[T]
	defer catch(&r.Err, nil)

	r.Val = fn()

	return r
}

func catch(cause *error, handler func(error)) {
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
