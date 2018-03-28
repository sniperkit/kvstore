package etcdv2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mickep76/kvstore"

	"github.com/coreos/etcd/client"
)

var (
	ErrNotSupported = errors.New("method is not supported by driver")
)

type conn struct {
	client *client.Client
}

func (c *conn) Close() error {
	return c.Close()
}

func (c *conn) Lease(ttl int) (kvstore.Lease, error) {
	return nil, ErrNotSupported
}

func (c *conn) Set(key string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("marshal value [%+v] for key [%s]: %v", value, key, err)
	}

	kvc := client.NewKeysAPI(*c.client)
	if _, err = kvc.Create(context.TODO(), key, string(b)); err != nil {
		return fmt.Errorf("set key [%s]: %v", key, err)
	}
	return nil
}

// TODO: pass lease as option to Set(key, value, opts...).
func (c *conn) SetWithLease(key string, value interface{}, l kvstore.Lease) error {
	return ErrNotSupported
}

func (c *conn) Delete(key string) error {
	kvc := client.NewKeysAPI(*c.client)
	if _, err := kvc.Delete(context.TODO(), key, &client.DeleteOptions{}); err != nil {
		return fmt.Errorf("delete key [%s]: %v", key, err)
	}
	return nil
}

func (c *conn) Keys(path string) ([]string, error) {
	/*
		resp, err := c.client.Get(context.TODO(), path, clientv3.WithPrefix(), clientv3.WithKeysOnly())
		if err != nil {
			return nil, fmt.Errorf("keys: %v", err)
		}

		keys := []string{}
		for _, kv := range resp.Kvs {
			keys = append(keys, string(kv.Key))
		}

		return keys, nil
	*/
	return nil, nil
}

func (c *conn) Values(key string) (kvstore.KeyValues, error) {
	/*
		resp, err := c.client.Get(context.TODO(), key, clientv3.WithPrefix())
		if err != nil {
			return nil, fmt.Errorf("values: %v", err)
		}

		values := kvstore.KeyValues{}
		for _, kv := range resp.Kvs {
			// TODO: add TTL for lease, if leaseID is 0 set nil for no lease.
			values = append(values, &kvstore.KeyValue{Key: string(kv.Key), Lease: &lease{id: clientv3.LeaseID(kv.Lease)}, Value: kvstore.Value(kv.Value)})
		}

		return values, nil
	*/
	return nil, nil
}

func (c *conn) Watch(path string) kvstore.Watch {
	/*
		return &watch{
			handlers: kvstore.WatchHandlers{},
			ch:       c.client.Watch(context.Background(), path, clientv3.WithPrefix()),
		}
	*/
	return nil
}
