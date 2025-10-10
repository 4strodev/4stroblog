package errors

import (
	"errors"
	"fmt"
)

type DomainError struct {
	code errorCode
	error
}

func (e *DomainError) Unwrap() error {
	return e.error
}

func WrapError(code errorCode, err error) *DomainError {
	domainError := new(DomainError)
	domainError.error = err
	domainError.code = code
	return domainError
}

func NewError(code errorCode, message string) *DomainError {
	err := errors.New(message)
	return WrapError(code, err)
}

func Errorf(code errorCode, message string, arguments ...any) *DomainError {
	err := fmt.Errorf(message, arguments...)
	return WrapError(code, err)
}

func (err *DomainError) Error() string {
	return err.error.Error()
}
