package cmp

import (
	"reflect"
	"regexp"
	"time"
)

// Support custom functions for compare of types

func Eq(a, b interface{}) (bool, error) {
	va := reflect.Indirect(reflect.ValueOf(a))
	vb := reflect.Indirect(reflect.ValueOf(b))

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

	case reflect.Struct:
		if va.Type().String() == "time.Time" && vb.Type().String() == "time.Time" {
			return va.Interface().(time.Time).Equal(vb.Interface().(time.Time)), nil
		}
		return false, ErrKindNotSupported
	}

	return false, ErrKindNotSupported
}

func Neq(a, b interface{}) (bool, error) {
	m, err := Eq(a, b)
	m = !m
	return m, err
}

func Lt(a, b interface{}) (bool, error) {
	va := reflect.Indirect(reflect.ValueOf(a))
	vb := reflect.Indirect(reflect.ValueOf(b))

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

	case reflect.Struct:
		if va.Type().String() == "time.Time" && vb.Type().String() == "time.Time" {
			return va.Interface().(time.Time).Before(vb.Interface().(time.Time)), nil
		}
		return false, ErrKindNotSupported
	}

	return false, ErrKindNotSupported
}

func Gt(a, b interface{}) (bool, error) {
	return Lt(b, a)
}

func Lte(a, b interface{}) (bool, error) {
	va := reflect.Indirect(reflect.ValueOf(a))
	vb := reflect.Indirect(reflect.ValueOf(b))

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

	case reflect.Struct:
		if va.Type().String() == "time.Time" && vb.Type().String() == "time.Time" {
			if va.Interface().(time.Time).Before(vb.Interface().(time.Time)) && va.Interface().(time.Time).Equal(vb.Interface().(time.Time)) {
				return true, nil
			}
			return false, nil
		}
		return false, ErrKindNotSupported
	}

	return false, ErrKindNotSupported
}

func Gte(a, b interface{}) (bool, error) {
	return Lte(b, a)
}

func Re(expr string, a interface{}) (bool, error) {
	re, err := regexp.Compile(expr)
	if err != nil {
		return false, err
	}

	va := reflect.Indirect(reflect.ValueOf(a))

	switch va.Kind() {
	case reflect.String:
		return re.MatchString(va.Interface().(string)), nil
	}

	return false, ErrKindNotSupported
}
