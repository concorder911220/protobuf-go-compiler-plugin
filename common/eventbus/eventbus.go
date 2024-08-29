package eventbus

import (
	"sync"
)

type EventBus struct {
	listeners map[string][]func(interface{}) interface{}
	mu        sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		listeners: make(map[string][]func(interface{}) interface{}),
	}
}

func (eb *EventBus) Subscribe(event string, callback func(interface{}) interface{}) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.listeners[event] = append(eb.listeners[event], callback)
}

func (eb *EventBus) Publish(event string, data interface{}) interface{} {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	if callbacks, ok := eb.listeners[event]; ok {
		for _, callback := range callbacks {
			return callback(data)
		}
	}
	return nil
}
