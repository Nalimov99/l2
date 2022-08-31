package handlers

import (
	"context"
	"dev11/internal/event"
	"dev11/internal/platform/web"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	ErrInvalidUserID = web.Error{
		Err:    errors.New("invalid user_id"),
		Status: http.StatusBadRequest,
	}
	ErrInvalidDate = web.Error{
		Err:    errors.New("invalid date"),
		Status: http.StatusBadRequest,
	}
	ErrInvalidEventID = web.Error{
		Err:    errors.New("invalid event_id"),
		Status: http.StatusBadRequest,
	}
	ErrEventWasNotFound = web.Error{
		Err:    errors.New("event was not found"),
		Status: http.StatusBadRequest,
	}
)

type Event struct {
	log      *log.Logger
	eventsFD *event.EventsFacade
}

func (e *Event) CreateEvent(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	var userID int
	ok, err := web.ParseRequestIntParam(r, "user_id", &userID)
	if err != nil || !ok || userID <= 0 {
		return &ErrInvalidUserID
	}

	var date time.Time
	_, err = web.ParseRequestDateParam(r, "date", &date)
	if err != nil {
		return &ErrInvalidDate
	}
	eventName := r.Form.Get("name")

	user := e.eventsFD.SetUserEvent(userID, date, eventName)

	return web.Respond(ctx, w, user, http.StatusOK)
}

func (e *Event) UpdateEvent(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	var eventID int
	ok, err := web.ParseRequestIntParam(r, "event_id", &eventID)
	if err != nil || !ok || eventID <= 0 {
		return &ErrInvalidEventID
	}

	eventName := r.Form.Get("name")
	var date time.Time
	hasDate := r.Form.Has("date")
	if hasDate {
		_, err = web.ParseRequestDateParam(r, "date", &date)
		if err != nil {
			return &ErrInvalidDate
		}
	}

	newEvent := event.UpdateEvent{
		ID: eventID,
		EventName: event.HasValue[string]{
			Value: eventName,
			Has:   r.Form.Has("name"),
		},
		Date: event.HasValue[time.Time]{
			Value: date,
			Has:   hasDate,
		},
	}

	event := e.eventsFD.UpdateEvent(newEvent)
	if event == nil {
		return &ErrEventWasNotFound
	}

	return web.Respond(ctx, w, event, http.StatusOK)
}

func (e *Event) DeleteEvent(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	var eventID int
	ok, err := web.ParseRequestIntParam(r, "event_id", &eventID)
	if err != nil || !ok || eventID <= 0 {
		return &ErrInvalidEventID
	}

	e.eventsFD.DeleteEvent(eventID)

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

func (e *Event) EventsForDay(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID, date, err := eventsParserHelper(r)
	if err != nil {
		return err
	}

	events := e.eventsFD.EventsForDuration(userID, date, 1, 0)

	return web.Respond(ctx, w, web.CommonRespond[[]*event.Event]{Result: events}, http.StatusOK)
}

func (e *Event) EventsForWeek(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID, date, err := eventsParserHelper(r)
	if err != nil {
		return err
	}

	events := e.eventsFD.EventsForDuration(userID, date, 8, 0)

	return web.Respond(ctx, w, web.CommonRespond[[]*event.Event]{Result: events}, http.StatusOK)
}

func (e *Event) EventsForMonth(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	userID, date, err := eventsParserHelper(r)
	if err != nil {
		return err
	}

	events := e.eventsFD.EventsForDuration(userID, date, 0, 1)

	return web.Respond(ctx, w, web.CommonRespond[[]*event.Event]{Result: events}, http.StatusOK)
}

func eventsParserHelper(r *http.Request) (userID int, date time.Time, err error) {
	m, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return userID, date, err
	}
	rDate, ok := m["date"]
	if !ok {
		return userID, date, &ErrInvalidDate
	}
	rUserID, ok := m["user_id"]
	if !ok {
		return userID, date, &ErrInvalidUserID
	}

	userID, err = strconv.Atoi(rUserID[0])
	if err != nil {
		return userID, date, &ErrInvalidUserID
	}
	date, err = time.Parse("2006-01-02", rDate[0])
	if err != nil {
		return userID, date, &ErrInvalidDate
	}

	return userID, date, nil
}
