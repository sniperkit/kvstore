package query

import (
	"reflect"
	"strings"
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
