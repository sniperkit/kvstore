package kvstore

import "errors"

var (
	ErrNotSupported = errors.New("method is not supported by driver")
	ErrEmptyValue   = errors.New("empty value")
	ErrNotSlice     = errors.New("not a slice")
	ErrNotStruct    = errors.New("not a struct")
)
