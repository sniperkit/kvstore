package etcdv3

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mickep76/kvstore"

	"github.com/coreos/etcd/clientv3"
)

type conn struct {
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

func (c *conn) Set(key string, value interface{}, options ...func(kvstore.KeyValue)) error {
	b, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("marshal value [%+v] for key [%s]: %v", value, key, err)
	}

	kvc := clientv3.NewKV(c.client)
	if _, err := kvc.Put(context.TODO(), key, string(b)); err != nil {
		return fmt.Errorf("set key [%s]: %v", key, err)
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
		values = append(values, &keyValue{key: string(kv.Key), lease: &lease{id: clientv3.LeaseID(kv.Lease)}, value: kvstore.Value(kv.Value)})
	}

	return values, nil
}

func (c *conn) Watch(path string) kvstore.Watch {
	return &watch{
		handlers: kvstore.WatchHandlers{},
		ch:       c.client.Watch(context.Background(), path, clientv3.WithPrefix()),
	}
}
