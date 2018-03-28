package etcdv2

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/mickep76/kvstore"

	"github.com/coreos/etcd/client"
)

type driver struct {
	timeout int
	tls     *tls.Config
	client.Config
}

func (d *driver) SetTimeout(timeout int) error {
	d.timeout = timeout
	return nil
}

func (d *driver) SetTLS(config *tls.Config) error {
	d.tls = config
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

	tr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   time.Duration(d.timeout) * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	if d.tls != nil {
		tr.TLSClientConfig = d.tls
	}

	c, err := client.New(d.Config)
	if err != nil {
		return nil, fmt.Errorf("conn: %v", err)
	}

	return &conn{
		client: &c,
	}, nil
}

func init() {
	kvstore.Register("etcdv2", &driver{
		timeout: 30,
	})
}
