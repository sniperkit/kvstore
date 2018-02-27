package kvstore

var drivers = make(map[string]Driver)

// Driver interface.
type Driver interface {
	SetTimeout(timeout int) error
	Open(endpoints []string) (Conn, error)
}

// Register database driver.
func Register(name string, driver Driver) {
	drivers[name] = driver
}
