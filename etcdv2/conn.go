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
	kvc := client.NewKeysAPI(*c.client)
	resp, err := kvc.Get(context.TODO(), path, nil)
	if err != nil {
		return nil, err
	}

	// TODO: add recursion and only fetch keys
	keys := []string{}
	for _, kv := range resp.Node.Nodes {
		keys = append(keys, string(kv.Key))
	}
	return keys, nil
}

func (c *conn) Values(key string) (kvstore.KeyValues, error) {
	kvc := client.NewKeysAPI(*c.client)
	resp, err := kvc.Get(context.TODO(), key, nil)
	if err != nil {
		return nil, err
	}

	// TODO: add recursion
	values := kvstore.KeyValues{}
	for _, kv := range resp.Node.Nodes {
		// TODO: add TTL
		values = append(values, &kvstore.KeyValue{Key: string(kv.Key), Value: kvstore.Value(kv.Value)})
	}
	return values, nil
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
