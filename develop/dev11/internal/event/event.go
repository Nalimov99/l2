package event

import (
	"sync"
	"time"
)

// NewEventsStore знает как заинициализировать пустое
// хранилище для мероприятий
func NewEventsStore() *EventsStore {
	return &EventsStore{
		mu:        sync.RWMutex{},
		LastIndex: 1,
		Store:     map[int]Event{},
	}
}

// Event возвращает мероприятие из хранилища.
// Если мероприятие не найдено, то вторым значением возвращает false
func (es *EventsStore) Event(id int) (*Event, bool) {
	es.mu.RLock()
	defer es.mu.Unlock()

	val, ok := es.Store[id]

	return &val, ok
}

// SetEvent знает как сохранить мероприятие в хранилище.
// Если в меропритие уже существует в хранилище, то данные старого мероприятия затруться
// Возвращаемым значением являеться ID созданного мероприятия
func (es *EventsStore) SetEvent(event Event) int {
	es.mu.Lock()
	defer es.mu.Unlock()

	currentIdx := es.LastIndex

	es.Store[currentIdx] = event
	es.LastIndex++

	return currentIdx
}

// DeleteEvent знает как удалить мероприятие из хранилища
func (es *EventsStore) DeleteEvent(id int) {
	es.mu.Lock()
	defer es.mu.Unlock()

	delete(es.Store, id)
}

// NewEvent знает как инициализировать мероприятие
func NewEvent(name string, date time.Time) *Event {
	return &Event{
		mu:   &sync.RWMutex{},
		Name: name,
		Date: date,
	}
}

// UpdateEvent знает как обновить данные мероприятия
func (e *Event) UpdateEvent(name string, date time.Time) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.Date = date
	e.Name = name
}
