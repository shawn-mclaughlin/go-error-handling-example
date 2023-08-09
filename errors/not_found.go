package errors

import (
	"fmt"
)

type notFound interface {
	NotFound() bool
}

type notFoundError struct {
	error
}

func (r notFoundError) NotFound() bool {
	return true
}

// IsNotFound returns true if err represents a missing resource
func IsNotFound(err error) bool {
	nf, ok := coreErrorAs[notFound](err)
	return ok && nf.NotFound()
}

// NewNotFoundError creates an error that represents a missing resource with the given message
func NewNotFoundError(msg string) error {
	return notFoundError{
		error: fmt.Errorf("not found: %s", msg),
	}
}

// WrapAsNotFoundError creates an error that represents a missing resource wrapping the given error
func WrapAsNotFoundError(err error) error {
	return notFoundError{
		error: fmt.Errorf("not found: %w", err),
	}
}
