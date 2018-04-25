package kvstore

import "crypto/tls"

var drivers = make(map[string]Driver)

// Driver interface.
type Driver interface {
	SetTimeout(timeout int) error
	SetTLS(config *tls.Config) error
	SetUser(user string) error
	SetPassword(password string) error
	SetEncoding(encoding string) error
	SetPrefix(prefix string) error
	Open(endpoints []string) (Conn, error)
}

// Register database driver.
func Register(name string, driver Driver) {
	drivers[name] = driver
}
