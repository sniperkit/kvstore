package kvstore

import "time"

type Lease interface {
	Renew() error
	KeepAlive() (chan *KeepAlive, error)
	StopKeepAlive() error
	Terminate() error
}

type KeepAlive struct {
	Timestamp time.Time
	Error     error
}
