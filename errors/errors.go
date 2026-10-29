// Package apierrors defines stable machine-readable error codes shared by the API.
package apierrors

import (
	"errors"
	"fmt"
)

// Code is a stable machine-readable error code.
type Code string

const (
	CodeNotFound      Code = "not_found"
	CodeAlreadyExists Code = "already_exists"
	CodeConflict      Code = "conflict"
	CodeInvalid       Code = "invalid_argument"
	CodeDenied        Code = "permission_denied"
	CodeUnavailable   Code = "unavailable"
	CodeInternal      Code = "internal"
)

// Error is the shared API error shape.
type Error struct {
	Code      Code
	Message   string
	Retryable bool
	Operation string
}

func (e *Error) Error() string {
	if e.Operation == "" {
		return fmt.Sprintf("%s: %s", e.Code, e.Message)
	}
	return fmt.Sprintf("%s: %s (operation %s)", e.Code, e.Message, e.Operation)
}

// New creates a typed API error.
func New(code Code, message string) *Error {
	return &Error{Code: code, Message: message}
}

// Is reports whether target is an *Error with the same code.
func Is(err error, code Code) bool {
	var target *Error
	if !errors.As(err, &target) {
		return false
	}
	return target.Code == code
}
