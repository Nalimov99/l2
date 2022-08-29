package handlers

import (
	"dev11/internal/event"
	"log"
	"net/http"
)

type Event struct {
	log    *log.Logger
	users  *event.UserStore
	events *event.EventsStore
}

func (e *Event) CreateEvent(w http.ResponseWriter, r *http.Request) {
	// body, _ := ioutil.ReadAll(r.Body)
	// log.Print(string(body))

	log.Print(r.URL.Query())
}
