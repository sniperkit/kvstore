package etcdv3

import (
	"github.com/mickep76/kvstore"

	"github.com/coreos/etcd/clientv3"
)

type watch struct {
	filter   *kvstore.EventType
	handlers kvstore.WatchHandlers
	ch       clientv3.WatchChan
}

func (w *watch) EventType(filter kvstore.EventType) kvstore.Watch {
	w.filter = &filter
	return w
}

func (w *watch) AddHandler(handler kvstore.WatchHandler) kvstore.Watch {
	w.handlers = append(w.handlers, handler)
	return w
}

func (w *watch) Start() error {
	for resp := range w.ch {
		for _, event := range resp.Events {
			kv := keyValue{key: string(event.Kv.Key), value: kvstore.Value(event.Kv.Value)}
			if event.IsCreate() {
				kv.event = &kvstore.Event{Type: kvstore.EventTypeCreate}
			} else if event.IsModify() {
				kv.event = &kvstore.Event{Type: kvstore.EventTypeModify}
			} else {
				kv.event = &kvstore.Event{Type: kvstore.EventTypeDelete}
			}

			if w.filter != nil && kv.event.Type != *w.filter {
				continue
			}

			for _, handler := range w.handlers {
				handler(kv)
			}
		}
	}
	return nil
}

func (w *watch) Stop() error {
	return nil
}
