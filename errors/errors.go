package errors

import (
	"errors"
)

type Error interface { }
// type Error interface {
// 	error
// }

type RouterError struct {
	Err 	error
}

func (re RouterError) Error() string {
	return re.Err.Error()
}

type StatusError struct {
	Err  	error
	Code 	int
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

func NewStatusError(err string, code int) StatusError {
	return StatusError{Err: errors.New(err), Code: code}
}

func NewRouterError(err string) RouterError {
	return RouterError{ Err: errors.New(err)}
}
