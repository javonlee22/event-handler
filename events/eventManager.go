package events

import (
	"fmt"
	"sync"
)

type eventMap = map[string]event

type eventManager struct {
	events eventMap
}

var manager *eventManager
var once sync.Once

func (em *eventManager) Subscribe(eventName string, subscriberID string, callback func(interface{})) {
	if val, ok := em.events[eventName]; ok {
		val.subscribe(subscriberID, callback)
	} else {
		e, _ := new(eventName, subscriberID, callback)
		em.events[eventName] = e
	}
}

func (em *eventManager) Unsubscribe(eventName string, subscriberID string) error {
	if val, ok := em.events[eventName]; ok {
		if err := val.unsubscribe(subscriberID); err != nil {
			return err
		}
	}
	return fmt.Errorf("Event with name  \"%s\" does not exist", eventName)
}

func (em eventManager) Emit(eventName string, data interface{}) error {
	if val, ok := em.events[eventName]; ok {
		val.emit(data)
	}
	return nil
}

// GetEventManager returns a singleton eventManager instance
func GetEventManager() *eventManager {
	once.Do(func() {
		manager = &eventManager{
			make(eventMap),
		}
	})
	return manager
}
