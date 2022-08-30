package event

import (
	"time"
)

// EventsFacade holds methods for UsersStore and EventsStore.
type EventsFacade struct {
	Users  *UserStore
	Events *EventsStore
}

// NewEventsFacade knows how to construct internal state for an EventsFacade
func NewEventsFacade() *EventsFacade {
	return &EventsFacade{
		Users:  NewUsersStore(),
		Events: NewEventsStore(),
	}
}

// SetUserEvent knows how to set Event into User and EventsStore.
// It will create new user if provided userID was not found in the store.
func (ef *EventsFacade) SetUserEvent(userID int, date time.Time, eventName string) *User {
	user, ok := ef.Users.User(userID)
	if !ok {
		user = ef.Users.SetUser(NewUser(userID))
	}

	event := NewEvent(eventName, date)
	ef.Events.SetEvent(event)

	user.SetUserEvent(event)
	return user
}

// UpdateEvent knows how to update event in the EventsStore
// It will return nil if event was not found in the store
func (ef *EventsFacade) UpdateEvent(updateEvent UpdateEvent) *Event {
	event, ok := ef.Events.Event(updateEvent.ID)
	if !ok {
		return nil
	}

	event.mu.Lock()
	defer event.mu.Unlock()
	if updateEvent.Date.Has {
		event.Date = updateEvent.Date.Value
	}
	if updateEvent.EventName.Has {
		event.Name = updateEvent.EventName.Value
	}

	ef.Events.SetEventByID(event, updateEvent.ID)

	return event
}

func (ef *EventsFacade) DeleteEvent(id int) {
	ef.Events.DeleteEvent(id)
}
