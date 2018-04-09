package kvstore

// Value of key.
type Value []byte

// KeyValue interface.
type KeyValue interface {
	Key() string
	Value() Value
	Decode(value interface{}) error
	Lease() Lease
	TTL() int
	Event() *Event
	SetLease(lease Lease) error
	SetTTL(ttl int) error
	SetEncoding(encoding string) error
}

// KeyValues multiple key/values.
type KeyValues []KeyValue
