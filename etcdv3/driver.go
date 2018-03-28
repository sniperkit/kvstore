package etcd

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/mickep76/kvstore"

	"github.com/coreos/etcd/clientv3"
)

type driver struct {
	clientv3.Config
}

func (d *driver) SetTimeout(timeout int) error {
	d.DialTimeout = time.Duration(timeout) * time.Second
	return nil
}

func (d *driver) SetTLS(config *tls.Config) error {
	d.TLS = config
	return nil
}

func (d *driver) SetUser(user string) error {
	d.Username = user
	return nil
}

func (d *driver) SetPassword(password string) error {
	d.Password = password
	return nil
}

func (d *driver) Open(endpoints []string) (kvstore.Conn, error) {
	d.Endpoints = endpoints

	c, err := clientv3.New(d.Config)
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
