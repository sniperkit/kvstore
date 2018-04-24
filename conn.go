package kvstore

import (
	"crypto/tls"
	"fmt"
)

// Conn connection interface.
type Conn interface {
	Close() error
	Lease(ttl int) (Lease, error)
	Set(key string, value interface{}, options ...func(KeyValue) error) error
	Delete(key string) error
	Keys(path string) ([]string, error)
	//	Value(path string) (KeyValue, error)
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

// WithTimeout for database connection.
func WithTimeout(timeout int) func(Driver) error {
	return func(d Driver) error {
		return d.SetTimeout(timeout)
	}
}

// WithTLS config for database connection.
func WithTLS(config *tls.Config) func(Driver) error {
	return func(d Driver) error {
		return d.SetTLS(config)
	}
}

// WithUser for database connection.
func WithUser(user string) func(Driver) error {
	return func(d Driver) error {
		return d.SetUser(user)
	}
}

// WithPassword for database connection.
func WithPassword(password string) func(Driver) error {
	return func(d Driver) error {
		return d.SetPassword(password)
	}
}

// WithEncoding to encode values.
func WithEncoding(encoding string) func(Driver) error {
	return func(d Driver) error {
		return d.SetEncoding(encoding)
	}
}

// WithLease for key/value.
func WithLease(lease Lease) func(KeyValue) error {
	return func(kv KeyValue) error {
		return kv.SetLease(lease)
	}
}

// WithTTL for key/value.
func WithTTL(ttl int) func(KeyValue) error {
	return func(kv KeyValue) error {
		return kv.SetTTL(ttl)
	}
}
