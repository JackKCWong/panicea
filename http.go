package panicea

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	StatusCode int
	Cause      error
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("status=%d, %s", e.StatusCode, e.Cause.Error())
}

func new(code int, err error, args ...interface{}) *HttpError {
	fmsg := args[0].(string)
	return &HttpError{code, fmt.Errorf(fmsg, err)}
}

func Unauthorized(err error, args ...interface{}) error {
	return new(http.StatusUnauthorized, err, args...)
}

func BadRequest(err error, args ...interface{}) error {
	return new(http.StatusBadRequest, err, args...)
}

func ServiceUnavailable(err error, args ...interface{}) error {
	return new(http.StatusServiceUnavailable, err, args...)
}

func InternalServerError(err error, args ...interface{}) error {
	return new(http.StatusInternalServerError, err, args...)
}
