package event

import (
	"sync"
	"time"
)

// User represents someone with Events
type User struct {
	mu     sync.RWMutex
	ID     int      `json:"id"`
	Events []*Event `json:"events"`
}

// UserStore contains all known users
type UserStore struct {
	Store map[int]*User
	mu    sync.RWMutex
}

// Event represents event's data
type Event struct {
	mu      sync.RWMutex
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	EventID int       `json:"event_id"`
	UserID  int       `json:"-"`
}

// EventsStore contains all known Events
type EventsStore struct {
	LastIndex int
	Store     map[int]*Event
	mu        sync.RWMutex
}

// UpdateEvent contains information needed to upate the event
type UpdateEvent struct {
	ID        int
	EventName HasValue[string]
	Date      HasValue[time.Time]
}

// HasValue should use to define that the provided value was passed
type HasValue[T any] struct {
	Value T
	Has   bool
}
