package errs

import "errors"

var (
	ErrRepository   = errors.New("repository error")
	ErrCustNotFound = errors.New("customer not found")
)
