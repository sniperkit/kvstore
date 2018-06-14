package kvstore

import (
	"crypto/tls"
	"fmt"
)

// Conn connection interface.
type Conn interface {
	Close() error
	Lease(ttl int) (Lease, error)
	Set(key string, value interface{}, options ...KeyValueOption) error
	Delete(key string) error
	Keys(path string) ([]string, error)
	//	Value(path string) (KeyValue, error)
	Values(path string) (KeyValues, error)
	Watch(path string) Watch
}

// DriverOption option for constructor.
type DriverOption func(Driver) error

// Open connection to database.
func Open(driver string, endpoints []string, options ...DriverOption) (Conn, error) {
	d, ok := drivers[driver]
	if !ok {
		return nil, fmt.Errorf("driver is not registered: %s", driver)
	}

	for _, option := range options {
		if err := option(d); err != nil {
			return nil, err
		}
	}

	return d.Open(endpoints)
}

// WithTimeout for database connection.
func WithTimeout(timeout int) DriverOption {
	return func(d Driver) error {
		return d.SetTimeout(timeout)
	}
}

// WithTLS config for database connection.
func WithTLS(config *tls.Config) DriverOption {
	return func(d Driver) error {
		return d.SetTLS(config)
	}
}

// WithUser for database connection.
func WithUser(user string) DriverOption {
	return func(d Driver) error {
		return d.SetUser(user)
	}
}

// WithPassword for database connection.
func WithPassword(password string) DriverOption {
	return func(d Driver) error {
		return d.SetPassword(password)
	}
}

// WithEncoding to encode values.
func WithEncoding(encoding string) DriverOption {
	return func(d Driver) error {
		return d.SetEncoding(encoding)
	}
}

// WithPrefix prefix pre+pended to key path.
func WithPrefix(prefix string) DriverOption {
	return func(d Driver) error {
		return d.SetPrefix(prefix)
	}
}

type KeyValueOption func(KeyValue) error

// WithLease for key/value.
func WithLease(lease Lease) KeyValueOption {
	return func(kv KeyValue) error {
		return kv.SetLease(lease)
	}
}

// WithTTL for key/value.
func WithTTL(ttl int) KeyValueOption {
	return func(kv KeyValue) error {
		return kv.SetTTL(ttl)
	}
}
