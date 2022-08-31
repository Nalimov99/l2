package event_test

import (
	"dev11/internal/event"
	"testing"
	"time"
)

func TestEvents(t *testing.T) {
	et := EventsTest{
		EF: event.NewEventsFacade(),
	}

	t.Log("RUN EVENTS FACADE TESTS")
	t.Run("EMPTY FACADE", et.EmptyFD)
	t.Run("ADD EVENT", et.AddEvent)
	t.Run("UPDATE EVENT", et.UpdateEvent)
	t.Run("DELETE EVENT", et.DeleteEvent)
	t.Run("EVENTS DURATION", et.DeleteEvent)
}

// EventsTest holds methods for each product subtest
// These type allows passing dependencies for tests
type EventsTest struct {
	EF  *event.EventsFacade
	Now time.Time
}

func (e *EventsTest) EmptyFD(t *testing.T) {
	if len(e.EF.Users.Users()) != 0 {
		t.Fatalf("Users should be empty, got:%d", len(e.EF.Users.Users()))
	}

	if len(e.EF.Events.Events()) != 0 {
		t.Fatalf("Events should be empty, got:%d", len(e.EF.Events.Events()))
	}
}

func (e *EventsTest) AddEvent(t *testing.T) {
	now := time.Now()

	e.Now = now

	user := e.EF.SetUserEvent(1, now, "test")

	if user.ID != 1 {
		t.Fatalf("user ID should be 1, got: %d", user.ID)
	}
	if len(user.Events) != 1 {
		t.Fatalf("user events should be 1, got: %d", len(user.Events))
	}
	event := user.Events[0]
	if event.Name != "test" {
		t.Fatalf("user event name should be 'test', got: %s", event.Name)
	}
}

func (e *EventsTest) UpdateEvent(t *testing.T) {
	now := time.Now()

	e.Now = now

	e.EF.UpdateEvent(event.UpdateEvent{
		ID: 1,
		EventName: event.HasValue[string]{
			Value: "test2",
			Has:   true,
		},
		Date: event.HasValue[time.Time]{
			Value: now,
			Has:   true,
		},
	})

	event := e.EF.Events.Store[1]

	if event.Name != "test2" {
		t.Fatalf("event name should be 'test2', got: %s", event.Name)
	}
	if event.Date != now {
		t.Fatalf("even date should be %v, got: %v", now, event.Date)
	}
	if event.UserID != 1 {
		t.Fatalf("event UserID should be 1, got: %d", event.UserID)
	}
	if event.EventID != 1 {
		t.Fatalf("event EventID should be 1, got: %d", event.EventID)
	}

	user, ok := e.EF.Users.User(1)
	if !ok {
		t.Fatalf("user with id 1 not found")
	}
	event = user.Events[0]
	if event.Name != "test2" {
		t.Fatalf("event name should be 'test2', got: %s", event.Name)
	}
	if event.Date != now {
		t.Fatalf("even date should be %v, got: %v", now, event.Date)
	}
	if event.UserID != 1 {
		t.Fatalf("event UserID should be 1, got: %d", event.UserID)
	}
	if event.EventID != 1 {
		t.Fatalf("event EventID should be 1, got: %d", event.EventID)
	}
}

func (e *EventsTest) DeleteEvent(t *testing.T) {
	e.EF.DeleteEvent(1)

	user, ok := e.EF.Users.User(1)
	if !ok {
		t.Fatalf("user with id 1 not found")
	}

	if len(user.Events) != 0 {
		t.Fatalf("user events should be 0, got: %d", len(user.Events))
	}
	if len(e.EF.Events.Store) != 0 {
		t.Fatalf("events should be 0 length, got: %d", len(e.EF.Events.Store))
	}
}

func (e *EventsTest) DurationEvents(t *testing.T) {
	now := time.Now()

	e.Now = now

	date0, _ := time.Parse("2006-01-02", "2022-08-31")
	e.EF.SetUserEvent(1, date0, "test")
	date1, _ := time.Parse("2006-01-02", "2022-09-01")
	e.EF.SetUserEvent(1, date1, "test")
	date2, _ := time.Parse("2006-01-02", "2022-09-02")
	e.EF.SetUserEvent(1, date2, "test")
	date3, _ := time.Parse("2006-01-02", "2022-09-07")
	e.EF.SetUserEvent(1, date3, "test")
	date4, _ := time.Parse("2006-01-02", "2022-09-08")
	e.EF.SetUserEvent(1, date4, "test")
	date5, _ := time.Parse("2006-01-02", "2022-09-30")
	e.EF.SetUserEvent(1, date5, "test")
	date6, _ := time.Parse("2006-01-02", "2022-10-01")
	e.EF.SetUserEvent(1, date6, "test")

	eventsForDay := len(e.EF.EventsForDuration(1, date1, 1, 0))
	eventsForWeek := len(e.EF.EventsForDuration(1, date1, 8, 0))
	eventsForMonth := len(e.EF.EventsForDuration(1, date1, 0, 1))

	if eventsForDay != 1 {
		t.Fatalf("events for day length should be 1, got: %d", eventsForDay)
	}

	if eventsForWeek != 3 {
		t.Fatalf("events for week length should be 3, got: %d", eventsForWeek)
	}

	if eventsForMonth != 5 {
		t.Fatalf("events for week length should be 5, got: %d", eventsForMonth)
	}
}
