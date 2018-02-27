package kvstore

// EventType type of event.
type EventType int

const (
	// EventTypeCreate create event.
	EventTypeCreate EventType = iota

	// EventTypeModify modify event.
	EventTypeModify

	// EventTypeDelete delete or expire event.
	EventTypeDelete
)

var eventTypeToString = map[EventType]string{
	EventTypeCreate: "create",
	EventTypeModify: "modify",
	EventTypeDelete: "delete",
}

// Event struct.
type Event struct {
	Type EventType
}

// String return event as a string.
func (e EventType) String() string {
	if v, ok := eventTypeToString[e]; ok {
		return v
	}
	return "unknown"
}
