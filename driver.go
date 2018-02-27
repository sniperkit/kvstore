package kvstore

var drivers = make(map[string]Driver)

type Driver interface {
	SetTimeout(timeout int) error
	Open(endpoints []string) (Conn, error)
}

func Register(name string, driver Driver) {
	drivers[name] = driver
}
