package kvstore

type EventType int

const (
	EventTypeCreate EventType = iota
	EventTypeModify
	EventTypeDelete
)

var eventTypeToString = map[EventType]string{
	EventTypeCreate: "create",
	EventTypeModify: "modify",
	EventTypeDelete: "delete",
}

type Event struct {
	Type EventType
}

func (e EventType) String() string {
	if v, ok := eventTypeToString[e]; ok {
		return v
	}
	return "unknown"
}
