package etcdv3

import (
	"github.com/mickep76/encdec"
	"github.com/mickep76/kvstore"
)

type keyValue struct {
	key      string
	lease    kvstore.Lease
	ttl      int
	prev     kvstore.Value
	value    kvstore.Value
	event    *kvstore.Event
	encoding string
}

type keyValues []*keyValue

func (kv keyValue) Key() string {
	return kv.key
}

func (kv keyValue) Value() kvstore.Value {
	return kv.value
}

func (kv keyValue) Lease() kvstore.Lease {
	return kv.lease
}

func (kv keyValue) Event() *kvstore.Event {
	return kv.event
}

func (kv keyValue) TTL() int {
	// TODO: if no lease return 0
	return kv.lease.TTL()
}

func (kv keyValue) SetLease(lease kvstore.Lease) error {
	kv.lease = lease
	return nil
}

func (kv keyValue) SetTTL(ttl int) error {
	return kvstore.ErrNotSupported
}

func (kv keyValue) SetEncoding(encoding string) error {
	if err := encdec.Registered(encoding); err != nil {
		return err
	}
	kv.encoding = encoding
	return nil
}
