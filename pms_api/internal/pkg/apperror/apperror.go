package apperror

import "errors"

type internalError struct {
	err error
}

func (e *internalError) Error() string {
	return e.err.Error()
}

var (
	NotFound = &internalError{err: errors.New("not found")}

	Unauthorized = &internalError{err: errors.New("unauthorized")}

	InvalidValue = &internalError{err: errors.New("invalid value")}
)
