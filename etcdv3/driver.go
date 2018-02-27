package etcd

import (
	"fmt"
	"time"

	"github.com/mickep76/kvstore"

	"github.com/coreos/etcd/clientv3"
)

type driver struct {
	timeout int
}

func (d *driver) SetTimeout(timeout int) error {
	d.timeout = timeout
	return nil
}

func (d *driver) Open(endpoints []string) (kvstore.Conn, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Duration(d.timeout) * time.Second,
	})
	if err != nil {
		return nil, fmt.Errorf("conn: %v", err)
	}

	return &conn{
		client: c,
	}, nil
}

func init() {
	kvstore.Register("etcdv3", &driver{})
}
