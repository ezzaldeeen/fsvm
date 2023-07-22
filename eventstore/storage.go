package eventstore

// Event ...
type Event interface {
	event()
}

// EventStore dummy event storage
type EventStore struct {
	kv map[string]Event
}

// NewEventStore factory function to instantiate EventStore
func NewEventStore(kv map[string]Event) *EventStore {
	return &EventStore{kv: kv}
}

// Save adds an event to the EventStore
func (e *EventStore) Save(key string, event Event) {
	e.kv[key] = event
}

// Load retrieves an event from the EventStore by key
func (e *EventStore) Load(key string) (Event, bool) {
	event, found := e.kv[key]
	return event, found
}
