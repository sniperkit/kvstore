package kvstore

type Value []byte

type KeyValue struct {
	Key string
	*Event
	Value
}

type KeyValues []*KeyValue

func (m Value) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
