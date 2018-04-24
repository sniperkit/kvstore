package etcdv3

import (
	"context"
	"fmt"

	"github.com/mickep76/kvstore"

	"github.com/coreos/etcd/clientv3"
	"github.com/mickep76/encdec"
)

type conn struct {
	encoding string

	client *clientv3.Client
}

func (c *conn) Close() error {
	return c.Close()
}

func (c *conn) Lease(ttl int) (kvstore.Lease, error) {
	if ttl < 5 {
		ttl = 5
	}

	r, err := c.client.Grant(context.TODO(), int64(ttl))
	if err != nil {
		return nil, fmt.Errorf("create lease: %v", err)
	}

	return &lease{
		id:     r.ID,
		ttl:    int64(ttl),
		client: c.client,
	}, nil
}

func (c *conn) Set(key string, value interface{}, options ...func(kvstore.KeyValue) error) error {
	kv := &keyValue{}
	for _, option := range options {
		if err := option(kv); err != nil {
			return err
		}
	}

	opts := []clientv3.OpOption{}
	if kv.lease != nil {
		opts = append(opts, clientv3.WithLease(kv.lease.(*lease).id))
	}

	kvc := clientv3.NewKV(c.client)

	// Check type
	if c.encoding == "" {
		switch value.(type) {
		case string:
			if _, err := kvc.Put(context.TODO(), key, value.(string), opts...); err != nil {
				return fmt.Errorf("set key [%s]: %v", key, err)
			}
		default:
			return fmt.Errorf("set key [%s]: value needs to be a string unless encoding is enabled", key)
		}
	} else {
		b, err := encdec.ToBytes(c.encoding, value)
		if err != nil {
			return err
		}

		if _, err := kvc.Put(context.TODO(), key, string(b), opts...); err != nil {
			return fmt.Errorf("set key [%s]: %v", key, err)
		}
	}
	return nil
}

func (c *conn) Delete(key string) error {
	kvc := clientv3.NewKV(c.client)
	if _, err := kvc.Delete(context.TODO(), key); err != nil {
		return fmt.Errorf("delete key [%s]: %v", key, err)
	}
	return nil
}

func (c *conn) Keys(path string) ([]string, error) {
	resp, err := c.client.Get(context.TODO(), path, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return nil, fmt.Errorf("keys: %v", err)
	}

	keys := []string{}
	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
	}

	return keys, nil
}

func (c *conn) Values(key string) (kvstore.KeyValues, error) {
	resp, err := c.client.Get(context.TODO(), key, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("values: %v", err)
	}

	values := kvstore.KeyValues{}
	for _, kv := range resp.Kvs {
		// TODO: add TTL for lease, if leaseID is 0 set nil for no lease.
		values = append(values, &keyValue{key: string(kv.Key), lease: &lease{id: clientv3.LeaseID(kv.Lease)}, value: kvstore.Value(kv.Value), encoding: c.encoding})
	}

	return values, nil
}

func (c *conn) Watch(path string) kvstore.Watch {
	return &watch{
		handlers: kvstore.WatchHandlers{},
		ch:       c.client.Watch(context.Background(), path, clientv3.WithPrefix()),
	}
}
