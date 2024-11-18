package pms_error

import "errors"

type internalError struct {
	err error
}

func (e *internalError) Error() string {
	return e.err.Error()
}

var (
	NotFound = &internalError{err: errors.New("not found")}
)
