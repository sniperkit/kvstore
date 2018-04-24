package etcdv3

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"

	"github.com/mickep76/kvstore"
)

type lease struct {
	id        clientv3.LeaseID
	ttl       int64
	client    *clientv3.Client
	keepAlive *time.Ticker
}

func (l *lease) TTL() int {
	return int(l.ttl)
}

func (l *lease) Renew() error {
	if _, err := l.client.KeepAliveOnce(context.TODO(), l.id); err != nil {
		return fmt.Errorf("renew lease: %v", err)
	}
	return nil
}

func (l *lease) KeepAlive() (chan *kvstore.KeepAlive, error) {
	ch := make(chan *kvstore.KeepAlive, 10)

	l.keepAlive = time.NewTicker(time.Duration(l.ttl/2) * time.Second)
	go func() {
		for t := range l.keepAlive.C {
			err := l.Renew()
			ch <- &kvstore.KeepAlive{
				Timestamp: t,
				Error:     err,
			}
		}
	}()

	return ch, nil
}

func (l *lease) StopKeepAlive() error {
	l.keepAlive.Stop()
	return nil
}

func (l *lease) Terminate() error {
	if _, err := l.client.Revoke(context.TODO(), l.id); err != nil {
		return fmt.Errorf("terminate lease: %v", err)
	}
	return nil
}
