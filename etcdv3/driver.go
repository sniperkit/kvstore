package etcdv3

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/mickep76/kvstore"

	"github.com/coreos/etcd/clientv3"
	"github.com/mickep76/encdec"
)

type driver struct {
	encoding string
	prefix   string

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

func (d *driver) SetEncoding(encoding string) error {
	if ok := encdec.Registered(encoding); !ok {
		return fmt.Errorf("encoding not registered: %s", encoding)
	}
	d.encoding = encoding
	return nil
}

func (d *driver) SetPrefix(prefix string) error {
	d.prefix = prefix
	return nil
}

func (d *driver) Open(endpoints []string) (kvstore.Conn, error) {
	d.Endpoints = endpoints

	c, err := clientv3.New(d.Config)
	if err != nil {
		return nil, fmt.Errorf("conn: %v", err)
	}

	return &conn{
		encoding: d.encoding,
		prefix:   d.prefix,
		client:   c,
	}, nil
}

func init() {
	kvstore.Register("etcdv3", &driver{})
}
