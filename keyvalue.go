package kvstore

import (
	"reflect"
)

// Value of key.
type Value []byte

// KeyValue interface.
type KeyValue interface {
	Key() string
	Value() Value
	PrevValue() Value
	Decode(value interface{}) error
	PrevDecode(value interface{}) error
	Lease() Lease
	TTL() int
	Event() *Event
	Encoding() string
	SetLease(lease Lease) error
	SetTTL(ttl int) error
}

// KeyValues multiple key/values.
type KeyValues []KeyValue

// Decode multiple values.
func (kvs KeyValues) Decode(value interface{}) error {
	p := reflect.ValueOf(value)
	if p.Kind() != reflect.Ptr {
		return ErrNotPtrListStruct
	}

	v := reflect.Indirect(p)
	if v.Kind() != reflect.Slice {
		return ErrNotPtrListStruct
	}

	for _, kv := range kvs {
		t := v.Type().Elem()
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}

		if t.Kind() != reflect.Struct {
			return ErrNotPtrListStruct
		}

		nv := reflect.New(t)

		if err := kv.Decode(nv.Interface()); err != nil {
			return err
		}

		v.Set(reflect.Append(v, nv))
	}

	return nil
}
