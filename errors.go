package panicea

import "fmt"


func Trap(fn func()) (reErr error) {
	defer catch(&reErr, nil)
	fn()

	return nil
}

func Try[T any](fn func() T) (val T, reErr error) {
	defer catch(&reErr, nil)

	return fn(), nil
}

func catch(cause *error, handler func(error)) {
	if caught := recover(); caught != nil {
		if err, ok := caught.(error); ok {
			*cause = err
			if handler != nil {
				handler(err)
			}
		} else {
			// propagate the original panic up
			panic(caught)
		}
	}
}

type Result[T any] struct {
	err error
	val T
}


func (r *Result[T]) On(msg string) T {
	if r.err != nil {
		panic(fmt.Errorf(msg, r.err))
	}

	return r.val
}

func Catch[T any](val T, err error) *Result[T] {
	return &Result[T]{
		err, val,
	}
}