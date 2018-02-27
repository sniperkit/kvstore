package kvstore

import "fmt"

// Conn connection interface.
type Conn interface {
	Close() error
	Lease(ttl int) (Lease, error)
	Set(key string, value interface{}) error
	SetWithLease(key string, value interface{}, lease Lease) error
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
