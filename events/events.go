package events

type EventsBus struct {
	subscribers map[string]map[string]func (Event) Response
}


func (e *EventsBus) subscribe(eventName, subName string, callback func (Event) Response) {
	e.subscribers[eventName][subName] = callback
}

func (e *EventsBus) publish(subType string, event Event) map[string]Response {
	responses := map[string]Response{}
	for subName, func := range e.subscribers[subType] {
		responses[subName] = func(event)
	}
	return responses
}

type Event struct {}

type Response struct {}
