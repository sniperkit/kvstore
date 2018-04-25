package kvstore

import "errors"

var (
	ErrNotSupported = errors.New("method is not supported by driver")
	ErrEmptyValue   = errors.New("empty value")
)