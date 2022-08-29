package event

import "sync"

// NewUserStore знает как заиницилазировать пусто UserStore
func NewUsersStore() *UserStore {
	users := make(map[int]User)

	return &UserStore{
		mu:    sync.RWMutex{},
		Store: users,
	}
}

// User возвращает пользователя из хранилища.
// Если пользователь не находиться в хранилище, то
// вторым значением вернет false.
func (us *UserStore) User(id int) (*User, bool) {
	us.mu.RLock()
	defer us.mu.RUnlock()

	val, ok := us.Store[id]

	return &val, ok
}

// SetUser знает как записать нового пользователя в хранилище.
// Перед вызовом SetUser, необходимо удостовериться что пользователь
// не записан в хранилище, иначе данные старого пользователя удаляться.
func (us *UserStore) SetUser(u User) {
	us.mu.Lock()
	defer us.mu.Unlock()

	us.Store[u.ID] = u
}

// NewUser знает как заиницилизировать пользователя с пустым списком мероприятий.
func NewUser(id int) User {
	return User{
		mu:     &sync.RWMutex{},
		ID:     id,
		Events: []int{},
	}
}
