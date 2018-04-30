package qry

import (
	"reflect"
	"strings"
)

func FieldValue(v interface{}, field string) (interface{}, error) {
	s := reflect.Indirect(reflect.ValueOf(v))
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
	s := reflect.Indirect(reflect.ValueOf(v))
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
