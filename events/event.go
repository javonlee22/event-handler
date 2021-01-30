package events

type subscribers = map[string]func(interface{})

type event struct {
	name        string
	subscribers subscribers
}

func new(name string, subscriberID string, callback func(interface{})) (event, error) {
	return event{
		name,
		subscribers{subscriberID: callback},
	}, nil
}

func (e event) getSubscriberCount() int {
	return len(e.subscribers)
}

func (e event) emit(data interface{}) error {
	for _, cb := range e.subscribers {
		cb(data)
	}
	return nil
}

func (e *event) subscribe(subscriberID string, callback func(interface{})) error {
	e.subscribers[subscriberID] = callback
	return nil
}

func (e *event) unsubscribe(subscriberID string) error {
	delete(e.subscribers, subscriberID)
	return nil
}
