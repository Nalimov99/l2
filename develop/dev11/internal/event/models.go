package event

import (
	"sync"
	"time"
)

// User структура содержащая данные о пользователе
// Events содержит набор мероприятий конкретного пользователя
type User struct {
	ID     int   `json:"id"`
	Events []int `json:"events"`
	mu     *sync.RWMutex
}

// UserStore знает все данные о всех известных пользователяй
type UserStore struct {
	Store map[int]User
	mu    sync.RWMutex
}

// Event данные мероприятия
type Event struct {
	Name string    `json:"name"`
	Date time.Time `json:"date"`
	mu   *sync.RWMutex
}

// EventsStore представляет из себя хранилище всех известных мероприятий
type EventsStore struct {
	LastIndex int
	Store     map[int]Event
	mu        sync.RWMutex
}
