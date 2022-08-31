package handlers

import (
	"dev11/internal/event"
	"dev11/internal/middleware"
	"dev11/internal/platform/web"
	"log"
	"net/http"
)

func API(log *log.Logger) http.Handler {
	app := web.NewApp(log, middleware.Logger(log), middleware.Errors(log))

	event := Event{
		log:      log,
		eventsFD: event.NewEventsFacade(),
	}

	app.Handle(http.MethodPost, "/create_event", event.CreateEvent)
	app.Handle(http.MethodPost, "/update_event", event.UpdateEvent)
	app.Handle(http.MethodPost, "/delete_event", event.DeleteEvent)
	app.Handle(http.MethodGet, "/events_for_day", event.EventsForDay)
	app.Handle(http.MethodGet, "/events_for_week", event.EventsForWeek)
	app.Handle(http.MethodGet, "/events_for_month", event.EventsForMonth)

	return app
}
