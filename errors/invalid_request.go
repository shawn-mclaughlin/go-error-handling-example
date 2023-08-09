package errors

import (
	"fmt"
)

type invalidRequest interface {
	InvalidRequest() bool
}

type invalidRequestError struct {
	error
}

func (r invalidRequestError) InvalidRequest() bool {
	return true
}

// IsInvalidRequest returns true if err represents a missing resource
func IsInvalidRequest(err error) bool {
	nf, ok := coreErrorAs[invalidRequest](err)
	return ok && nf.InvalidRequest()
}

// NewInvalidRequestError creates an error that represents receiving an invalid request adding the given message
func NewInvalidRequestError(msg string) error {
	return invalidRequestError{
		error: fmt.Errorf("not found: %s", msg),
	}
}

// WrapAsInvalidRequestError creates an error that represents receiving an invalid request wrapping the given error
func WrapAsInvalidRequestError(err error) error {
	return invalidRequestError{
		error: fmt.Errorf("not found: %w", err),
	}
}
