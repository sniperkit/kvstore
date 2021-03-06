package etcdv3

import (
	"github.com/mickep76/kvstore"

	"github.com/mickep76/encdec"
	_ "github.com/mickep76/encdec/json"
	_ "github.com/mickep76/encdec/toml"
	_ "github.com/mickep76/encdec/yaml"
)

type keyValue struct {
	key       string
	lease     kvstore.Lease
	ttl       int
	value     kvstore.Value
	prevValue kvstore.Value
	event     *kvstore.Event
	encoding  string
}

type keyValues []*keyValue

func (kv *keyValue) Key() string {
	return kv.key
}

func (kv *keyValue) Value() kvstore.Value {
	return kv.value
}

func (kv *keyValue) PrevValue() kvstore.Value {
	return kv.prevValue
}

func (kv *keyValue) Decode(value interface{}) error {
	if kv.value == nil {
		return kvstore.ErrEmptyValue
	}

	if err := encdec.FromBytes(kv.encoding, kv.value, value); err != nil {
		return err
	}
	return nil
}

func (kv *keyValue) PrevDecode(value interface{}) error {
	if kv.prevValue == nil {
		return kvstore.ErrEmptyValue
	}

	if err := encdec.FromBytes(kv.encoding, kv.prevValue, value); err != nil {
		return err
	}
	return nil
}

func (kv *keyValue) Lease() kvstore.Lease {
	return kv.lease
}

func (kv *keyValue) TTL() int {
	if kv.lease == nil {
		return 0
	}
	return kv.lease.TTL()
}

func (kv *keyValue) Event() *kvstore.Event {
	return kv.event
}

func (kv *keyValue) Encoding() string {
	return kv.encoding
}

func (kv *keyValue) SetLease(lease kvstore.Lease) error {
	kv.lease = lease
	return nil
}

func (kv *keyValue) SetTTL(ttl int) error {
	return kvstore.ErrNotSupported
}
