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
		return false, ErrNotSameKind
	}

	switch va.Kind() {
	case reflect.Bool:
		return a.(bool) == b.(bool), nil

	case reflect.Int:
		return a.(int) == b.(int), nil
	case reflect.Int8:
		return a.(int8) == b.(int8), nil
	case reflect.Int16:
		return a.(int16) == b.(int16), nil
	case reflect.Int32:
		return a.(int32) == b.(int32), nil
	case reflect.Int64:
		return a.(int64) == b.(int64), nil

	case reflect.Uint:
		return a.(uint) == b.(uint), nil
	case reflect.Uint8:
		return a.(uint8) == b.(uint8), nil
	case reflect.Uint16:
		return a.(uint16) == b.(uint16), nil
	case reflect.Uint32:
		return a.(uint32) == b.(uint32), nil
	case reflect.Uint64:
		return a.(uint64) == b.(uint64), nil

	case reflect.String:
		return a.(string) == b.(string), nil
	}

	return false, ErrKindNotSupported
}

func CmpNeq(a, b interface{}) (bool, error) {
	m, err := CmpEq(a, b)
	m = !m
	return m, err
}

func CmpLt(a, b interface{}) (bool, error) {
	va, err := resolvePtr(reflect.ValueOf(a))
	if err != nil {
		return false, err
	}

	vb, err := resolvePtr(reflect.ValueOf(b))
	if err != nil {
		return false, err
	}

	if va.Kind() != vb.Kind() {
		return false, ErrNotSameKind
	}

	switch va.Kind() {
	case reflect.Int:
		return a.(int) < b.(int), nil
	case reflect.Int8:
		return a.(int8) < b.(int8), nil
	case reflect.Int16:
		return a.(int16) < b.(int16), nil
	case reflect.Int32:
		return a.(int32) < b.(int32), nil
	case reflect.Int64:
		return a.(int64) < b.(int64), nil

	case reflect.Uint:
		return a.(uint) < b.(uint), nil
	case reflect.Uint8:
		return a.(uint8) < b.(uint8), nil
	case reflect.Uint16:
		return a.(uint16) < b.(uint16), nil
	case reflect.Uint32:
		return a.(uint32) < b.(uint32), nil
	case reflect.Uint64:
		return a.(uint64) < b.(uint64), nil

	case reflect.String:
		return a.(string) < b.(string), nil
	}

	return false, ErrKindNotSupported
}

func CmpGt(a, b interface{}) (bool, error) {
	return CmpLt(b, a)
}

func CmpLte(a, b interface{}) (bool, error) {
	va, err := resolvePtr(reflect.ValueOf(a))
	if err != nil {
		return false, err
	}

	vb, err := resolvePtr(reflect.ValueOf(b))
	if err != nil {
		return false, err
	}

	if va.Kind() != vb.Kind() {
		return false, ErrNotSameKind
	}

	switch va.Kind() {
	case reflect.Int:
		return a.(int) <= b.(int), nil
	case reflect.Int8:
		return a.(int8) <= b.(int8), nil
	case reflect.Int16:
		return a.(int16) <= b.(int16), nil
	case reflect.Int32:
		return a.(int32) <= b.(int32), nil
	case reflect.Int64:
		return a.(int64) <= b.(int64), nil

	case reflect.Uint:
		return a.(uint) <= b.(uint), nil
	case reflect.Uint8:
		return a.(uint8) <= b.(uint8), nil
	case reflect.Uint16:
		return a.(uint16) <= b.(uint16), nil
	case reflect.Uint32:
		return a.(uint32) <= b.(uint32), nil
	case reflect.Uint64:
		return a.(uint64) <= b.(uint64), nil

	case reflect.String:
		return a.(string) <= b.(string), nil
	}

	return false, ErrKindNotSupported
}

func CmpGte(a, b interface{}) (bool, error) {
	return CmpLte(b, a)
}
