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
			kv := keyValue{Key: string(event.Kv.Key), Value: kvstore.Value(event.Kv.Value)}
			if event.IsCreate() {
				kv.Event = &kvstore.Event{Type: kvstore.EventTypeCreate}
			} else if event.IsModify() {
				kv.Event = &kvstore.Event{Type: kvstore.EventTypeModify}
			} else {
				kv.Event = &kvstore.Event{Type: kvstore.EventTypeDelete}
			}

			if w.filter != nil && kv.Type != *w.filter {
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
