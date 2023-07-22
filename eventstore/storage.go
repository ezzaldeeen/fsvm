package eventstore

// Event ...
type Event interface {
	event()
}

type Events []Event

// EventStore dummy event storage
type EventStore struct {
	kv map[string]Events
}

// NewEventStore factory function to instantiate EventStore
func NewEventStore(kv map[string]Events) EventStore {
	return EventStore{kv: kv}
}

// Save adds an event to the EventStore
func (e *EventStore) Save(key string, events Events) {
	e.kv[key] = events
}

// Load retrieves an event from the EventStore by key
func (e *EventStore) Load(key string) (Events, bool) {
	events, found := e.kv[key]
	return events, found
}
