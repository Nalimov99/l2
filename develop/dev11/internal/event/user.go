package event

import (
	"sync"
)

// NewUserStore knows how to construct internal state for an NewUserStore
func NewUsersStore() *UserStore {
	users := make(map[int]*User)

	return &UserStore{
		mu:    sync.RWMutex{},
		Store: users,
	}
}

// User knows how to retrieve User from UsersStore
// If specified user was not exists in the store,
// it will return false in the second argument
func (us *UserStore) User(id int) (*User, bool) {
	us.mu.RLock()
	defer us.mu.RUnlock()

	val, ok := us.Store[id]

	return val, ok
}

// Users knows how to retrieve Store
func (us *UserStore) Users() map[int]*User {
	us.mu.RLock()
	defer us.mu.RUnlock()

	return us.Store
}

// SetUser knows how to set User into the UsersStore
func (us *UserStore) SetUser(u *User) *User {
	us.mu.Lock()
	defer us.mu.Unlock()

	us.Store[u.ID] = u

	return u
}

// NewUser knows how to construct internal state for an User
func NewUser(id int) *User {
	return &User{
		mu:     sync.RWMutex{},
		ID:     id,
		Events: []*Event{},
	}
}

func (u *User) SetUserEvent(event *Event) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.Events = append(u.Events, event)
}

func (us *UserStore) DeleteUserEvent(userID, eventID int) {
	user, ok := us.User(userID)
	if !ok {
		return
	}

	user.mu.RLock()
	var elemIdx int

	for i, event := range user.Events {
		if event.EventID == eventID {
			elemIdx = i
			break
		}
	}

	newEvents := make([]*Event, 0)
	newEvents = append(newEvents, user.Events[:elemIdx]...)
	newEvents = append(newEvents, user.Events[elemIdx+1:]...)
	user.mu.RUnlock()
	user.mu.Lock()
	user.Events = newEvents
	user.mu.Unlock()
}
