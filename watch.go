package kvstore

type (
	// WatchHandler handler for watch events.
	WatchHandler func(KeyValue)

	// WatchHandlers list of handlers for watch events.
	WatchHandlers []WatchHandler
)

// Watch interface for intercepting create, modify or delete events to keys.
type Watch interface {
	// EventType filter on event type.
	EventType(filter EventType) Watch

	// AddHandler to a watcher.
	AddHandler(handler WatchHandler) Watch

	// Start watcher.
	Start() error

	// Stop watcher.
	Stop() error
}
