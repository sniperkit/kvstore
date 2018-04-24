package kvstore

import "time"

// Lease struct.
type Lease interface {
	// TTL for lease.
	TTL() int

	// Renew lease.
	Renew() error

	// KeepAlive lease periodically every (ttl/2).
	KeepAlive() (chan *KeepAlive, error)

	// StopKeepAlive stop renewing lease.
	StopKeepAlive() error

	// Terminate lease.
	Terminate() error
}

// KeepAlive struct.
type KeepAlive struct {
	Timestamp time.Time
	Error     error
}
