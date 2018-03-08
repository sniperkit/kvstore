package kvstore

// Value of key.
type Value []byte

// KeyValue struct contains event, key and value.
type KeyValue struct {
	Key   string
	Lease Lease
	*Event
	Value
}

// KeyValues multiple key/values.
type KeyValues []*KeyValue

// MarshalJSON for key/value.
func (m Value) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
