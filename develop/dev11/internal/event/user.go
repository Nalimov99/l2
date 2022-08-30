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
