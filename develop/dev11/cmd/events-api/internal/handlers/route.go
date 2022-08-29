package handlers

import (
	"dev11/internal/event"
	"log"
	"net/http"
)

func API(log *log.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	usersStore := event.NewUsersStore()
	eventsStore := event.NewEventsStore()

	event := Event{
		log:    log,
		users:  usersStore,
		events: eventsStore,
	}

	mux.HandleFunc("/create_event", event.CreateEvent)

	return mux
}
