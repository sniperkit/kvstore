package kvstore

// Value of key.
type Value []byte

// KeyValue interface.
type KeyValue interface {
	Key() string
	Value() Value
	PrevValue() Value
	Decode(value interface{}) error
	PrevDecode(value interface{}) error
	Lease() Lease
	TTL() int
	Event() *Event
	Encoding() string
	SetLease(lease Lease) error
	SetTTL(ttl int) error
}

// KeyValues multiple key/values.
type KeyValues []KeyValue
