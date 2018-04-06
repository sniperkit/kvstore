package kvstore

import (
	"crypto/tls"
	"fmt"
)

// Conn connection interface.
type Conn interface {
	Close() error
	Lease(ttl int) (Lease, error)
	Set(key string, value interface{}, options ...func(KeyValue)) error
	Delete(key string) error
	Keys(path string) ([]string, error)
	Values(path string) (KeyValues, error)
	Watch(path string) Watch
}

// Open connection to database.
func Open(driver string, endpoints []string, options ...func(Driver) error) (Conn, error) {
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

// Timeout for database connection.
func Timeout(timeout int) func(Driver) error {
	return func(d Driver) error {
		return d.SetTimeout(timeout)
	}
}

// TLS config for database connection.
func TLS(config *tls.Config) func(Driver) error {
	return func(d Driver) error {
		return d.SetTLS(config)
	}
}

// User for database connection.
func User(user string) func(Driver) error {
	return func(d Driver) error {
		return d.SetUser(user)
	}
}

// Password for database connection.
func Password(password string) func(Driver) error {
	return func(d Driver) error {
		return d.SetPassword(password)
	}
}

// OpLease for key/value.
func OpLease(lease Lease) func(KeyValue) error {
	return func(kv KeyValue) error {
		return kv.SetLease(lease)
	}
}

// OpTTL for key/value.
func OpTTL(ttl int) func(KeyValue) error {
	return func(kv KeyValue) error {
		return kv.SetTTL(ttl)
	}
}
