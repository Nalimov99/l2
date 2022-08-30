package handlers

import (
	"context"
	"dev11/internal/event"
	"dev11/internal/platform/web"
	"errors"
	"log"
	"net/http"
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

	userID, err := strconv.Atoi(r.Form.Get("user_id"))
	if err != nil || userID <= 0 {
		return &ErrInvalidUserID
	}

	date, err := time.Parse(time.RFC3339, r.Form.Get("date"))
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

	eventID, err := strconv.Atoi(r.Form.Get("event_id"))
	if err != nil || eventID <= 0 {
		return &ErrInvalidEventID
	}
	eventName := r.Form.Get("name")
	var date time.Time
	hasDate := r.Form.Has("date")
	if hasDate {
		date, err = time.Parse(time.RFC3339, r.Form.Get("date"))
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

	eventID, err := strconv.Atoi(r.Form.Get("event_id"))
	if err != nil || eventID <= 0 {
		return &ErrInvalidEventID
	}

	e.eventsFD.DeleteEvent(eventID)

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}
