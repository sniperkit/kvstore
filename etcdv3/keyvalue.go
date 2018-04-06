package etcdv3

import (
	"github.com/mickep76/kvstore"
)

type keyValue struct {
	Key   string
	Lease kvstore.Lease
	TTL   int
	Prev  kvstore.Value
	*kvstore.Event
	kvstore.Value
}

type keyValues []*keyValue

func (kv keyValue) SetLease(lease kvstore.Lease) error {
	kv.Lease = lease
	return nil
}

func (kv keyValue) SetTTL(ttl int) error {
	return kvstore.ErrNotSupported
}
