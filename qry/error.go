package qry

import (
	"errors"
)

var (
	ErrInvalidPtr       = errors.New("invalid ptr")
	ErrNotAStruct       = errors.New("not a struct")
	ErrUnknownField     = errors.New("unknown field")
	ErrNotSameKind      = errors.New("not same kind")
	ErrKindNotSupported = errors.New("kind not supported")
	ErrUnknownOperator  = errors.New("unknown operator")
)
