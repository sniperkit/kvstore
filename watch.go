package kvstore

type (
	WatchHandler  func(KeyValue)
	WatchHandlers []WatchHandler
)

type Watch interface {
	EventType(filter EventType) Watch
	AddHandler(handler WatchHandler) Watch
	Start() error
	Stop() error
}
