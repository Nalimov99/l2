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

// Events knows how to retrieve Store
func (es *EventsStore) Events() map[int]*Event {
	es.mu.RLock()
	defer es.mu.RUnlock()

	return es.Store
}

// SetEvent knows how to write new Event into the EventsStore
// Return value is id of created Event
func (es *EventsStore) SetEvent(event *Event, userID int) {
	es.mu.Lock()
	defer es.mu.Unlock()

	currentIdx := es.LastIndex
	event.EventID = currentIdx
	event.UserID = userID

	es.Store[currentIdx] = event
	es.LastIndex++
}

func (es *EventsStore) SetEventByID(event *Event, id int) {
	es.mu.Lock()
	defer es.mu.Unlock()

	es.Store[id] = event
}

// DeleteEvent knows how to delete Event from EventsStore
// It will return event_id and user_id of deleted elemnt
func (es *EventsStore) DeleteEvent(id int) (eventID int, userID int, ok bool) {
	event, ok := es.Event(id)
	if !ok {
		return eventID, userID, ok
	}

	es.mu.Lock()
	delete(es.Store, id)
	es.mu.Unlock()
	return event.EventID, event.UserID, true
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
