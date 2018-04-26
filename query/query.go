package query

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	ErrInvalidPtr    = errors.New("invalid ptr")
	ErrNotAStruct    = errors.New("not a struct")
	ErrUnknownField  = errors.New("unknown field")
	ErrDifferentKind = errors.New("different kind")
)

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

func GetFieldValue(v interface{}, field string) (interface{}, error) {
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

func GetFieldTagValue(v interface{}, field string, tag string) (interface{}, error) {
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

func CmpEq(a, b interface{}) (bool, error) {
	va, err := resolvePtr(reflect.ValueOf(a))
	if err != nil {
		return false, err
	}

	vb, err := resolvePtr(reflect.ValueOf(b))
	if err != nil {
		return false, err
	}

	if va.Kind() != vb.Kind() {
		return false, ErrDifferentKind
	}

	return fmt.Sprint(va.Interface()) == fmt.Sprint(vb.Interface()), nil
}

func CmpNeq(a, b interface{}) (bool, error) {
	m, err := CmpEq(a, b)
	m = !m
	return m, err
}

func CmpLt() (bool, error) {
	return false, nil
}

func CmpGt() (bool, error) {
	return CmpLt(b, a)
}
