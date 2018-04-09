package kvstore

// Value of key.
type Value []byte

// KeyValue interface.
type KeyValue interface {
	Key() string
	Value() Value
	Lease() Lease
	TTL() int
	Event() *Event
	SetLease(lease Lease) error
	SetTTL(ttl int) error
	SetEncoding(encoding string) error
}

// KeyValues multiple key/values.
type KeyValues []KeyValue

// MarshalJSON for key/value.
func (m Value) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
