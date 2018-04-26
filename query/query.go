package query

import (
	"errors"
	"reflect"
	"strings"
)

var (
	ErrInvalidPtr       = errors.New("invalid ptr")
	ErrNotAStruct       = errors.New("not a struct")
	ErrUnknownField     = errors.New("unknown field")
	ErrNotSameKind      = errors.New("not same kind")
	ErrKindNotSupported = errors.New("kind not supported")
)

type Operator string

const (
	GROUP Operator = "GROUP"
	AND   Operator = "AND"
	OR    Operator = "OR"
	NOT   Operator = "NOT"
	EQ    Operator = "EQ"
	NEQ   Operator = "NEQ"
	LT    Operator = "LT"
	LTE   Operator = "LTE"
	GT    Operator = "GT"
	GTE   Operator = "GTE"
	IN    Operator = "IN"
	RE    Operator = "RE"
)

type Query struct {
	Oper Operator `json:"operator,omitempty"`

	*QueryField
	*QueryMatch
	*QueryGroup
}

type Queries []*Query

type QueryField struct {
	Field string `json:"field"`
}

type QueryMatch struct {
	Value interface{} `json:"value,omitempty"`
}

type QueryGroup struct {
	Queries Queries `json:"queries"`
}

func resolvePtr(v reflect.Value) (reflect.Value, error) {
	if v.Kind() == reflect.Ptr {
		var err error

		v, err = resolvePtr(v.Elem())
		if err != nil {
			return reflect.Value{}, err
		}

		if !v.IsValid() {
			return reflect.Value{}, ErrInvalidPtr
		}
	}

	return v, nil
}

func FieldValue(v interface{}, field string) (interface{}, error) {
	s, err := resolvePtr(reflect.ValueOf(v))
	if err != nil {
		return nil, err
	}

	if s.Kind() != reflect.Struct {
		return nil, ErrNotAStruct
	}

	f := s.FieldByName(field)
	if !f.IsValid() {
		return nil, ErrUnknownField
	}

	return f.Interface(), nil
}

func FieldTagValue(v interface{}, field string, tag string) (interface{}, error) {
	s, err := resolvePtr(reflect.ValueOf(v))
	if err != nil {
		return nil, err
	}

	if s.Kind() != reflect.Struct {
		return nil, ErrNotAStruct
	}

	for i := 0; i < s.NumField(); i++ {
		f := s.Type().Field(i)
		t, ok := f.Tag.Lookup(tag)
		if !ok {
			continue
		}

		n := strings.Split(t, ",")[0]
		if field == n {
			return s.Field(i).Interface(), nil
		}
	}

	return nil, ErrUnknownField
}
