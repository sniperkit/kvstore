package etcdv2

import (
	"github.com/mickep76/kvstore"
)

type lease struct {
}

func (l *lease) Renew() error {
	return ErrNotSupported
}

func (l *lease) KeepAlive() (chan *kvstore.KeepAlive, error) {
	return nil, ErrNotSupported
}

func (l *lease) StopKeepAlive() error {
	return ErrNotSupported
}

func (l *lease) Terminate() error {
	return ErrNotSupported
}
