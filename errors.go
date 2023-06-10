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

func (r *Result[T]) Throw(args ...interface{}) T {
	if r.err == nil {
		return r.val
	}

	if len(args) == 1 {
		switch arg := args[0].(type) {
		case string:
			panic(fmt.Errorf(arg, r.err))
		case error:
			panic(arg)
		case func(error) error:
			panic(arg(r.err))
		}
	}

	if len(args) >= 2 {
		switch arg0 := args[0].(type) {
		case func(error, ...interface{}) error:
			panic(arg0(r.err, args[1:]...))
		}
	}

	panic(fmt.Errorf("unexpected args: %v", args))
}

func Catch[T any](val T, err error) *Result[T] {
	return &Result[T]{
		err, val,
	}
}
