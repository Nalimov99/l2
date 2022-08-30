package event

import (
	"sync"
	"time"
)

// NewEventsStore knows how to construct internal state for an EventsStore
func NewEventsStore() *EventsStore {
	return &EventsStore{
		mu:        sync.RWMutex{},
		LastIndex: 1,
		Store:     map[int]*Event{},
	}
}

// Event knows how to retrieve id from EventsStore.
// If the value was not found it will return false
func (es *EventsStore) Event(id int) (*Event, bool) {
	es.mu.RLock()
	defer es.mu.RUnlock()

	val, ok := es.Store[id]

	return val, ok
}

// SetEvent knows how to write new Event into the EventsStore
// Return value is id of created Event
func (es *EventsStore) SetEvent(event *Event) {
	es.mu.Lock()
	defer es.mu.Unlock()

	currentIdx := es.LastIndex

	es.Store[currentIdx] = event
	es.LastIndex++
}

func (es *EventsStore) SetEventByID(event *Event, id int) {
	es.mu.Lock()
	defer es.mu.Unlock()

	es.Store[id] = event
}

// DeleteEvent knows how to delete Event from EventsStore
func (es *EventsStore) DeleteEvent(id int) {
	es.mu.Lock()
	defer es.mu.Unlock()

	delete(es.Store, id)
}

// NewEvent  knows how to construct internal state for an Event
func NewEvent(name string, date time.Time) *Event {
	return &Event{
		mu:   sync.RWMutex{},
		Name: name,
		Date: date,
	}
}

// UpdateEvent knows how to update the Event in the EventsStore
func (e *Event) UpdateEvent(name string, date time.Time) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.Date = date
	e.Name = name
}
