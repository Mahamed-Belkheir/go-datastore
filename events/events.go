package events

type EventsBus struct {
	subscribers map[string]map[string]func(interface{}) interface{}
}

func (e *EventsBus) Subscribe(eventName, subName string, callback func(interface{}) interface{}) {
	e.subscribers[eventName][subName] = callback
}

func (e *EventsBus) Publish(subType string, event interface{}) map[string]interface{} {
	responses := map[string]interface{}{}
	for subName, callback := range e.subscribers[subType] {
		responses[subName] = callback(event)
	}
	return responses
}
