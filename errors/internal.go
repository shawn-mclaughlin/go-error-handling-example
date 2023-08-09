package errors

import "errors"

func coreErrorAs[E any](err error) (E, bool) {
	var target E
	ok := errors.As(err, &target)

	return target, ok
}
