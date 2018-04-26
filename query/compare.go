package query

import (
	"reflect"
	"regexp"
)

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

func CmpRe(expr string, a interface{}) (bool, error) {
	re, err := regexp.Compile(expr)
	if err != nil {
		return false, err
	}

	va, err := resolvePtr(reflect.ValueOf(a))
	if err != nil {
		return false, err
	}

	switch va.Kind() {
	case reflect.String:
		return re.MatchString(va.Interface().(string)), nil
	}

	return false, ErrKindNotSupported
}

func CmpIn(a, b interface{}) (bool, error) {
	return false, nil
}
