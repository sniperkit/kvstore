package qry

import (
	"errors"
)

var (
	ErrNotAStruct       = errors.New("not a struct")
	ErrUnknownField     = errors.New("unknown field")
	ErrNotSameKind      = errors.New("not same kind")
	ErrKindNotSupported = errors.New("kind not supported")
)
