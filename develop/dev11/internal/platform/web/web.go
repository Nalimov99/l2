package web

import (
	"context"
	"errors"
	"log"
	"net/http"
)

var (
	ErrContextValueMissing = errors.New("web values missing from context")
)

// ctxCommonValuesKey represents the type of value for the context key
type ctxCommonValuesKey string

// KeyValues is how request values or stored/retrieve
const KeyValues ctxCommonValuesKey = "commonValues"

// ContextValues carries information about each request
type ContexValues struct {
	StatusCode int
}

//Handler is a signature that all applications handlers will implement
type Handler func(context.Context, http.ResponseWriter, *http.Request) error

//App is the entry point for all web aplications
type App struct {
	mux *http.ServeMux
	log *log.Logger
	mw  []Middleware
}

//NewApp knows how to construct internal state for an App
func NewApp(log *log.Logger, mw ...Middleware) *App {
	return &App{
		mux: http.NewServeMux(),
		log: log,
		mw:  mw,
	}
}

//Handle connects a method and URL pattern to a particular application handler
func (a *App) Handle(method, pattern string, h Handler) {

	// Add aplications general middleware
	h = wrapMiddleware(a.mw, h)

	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			log.Printf(
				"METHOD NOT ALLOWED: %d %s %s",
				http.StatusMethodNotAllowed, r.Method, r.URL.Path,
			)
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		v := ContexValues{}

		ctx := context.WithValue(r.Context(), KeyValues, &v)

		if err := h(ctx, w, r); err != nil {
			a.log.Printf("ERROR: Unhandled error: %v", err)
		}
	}

	a.mux.HandleFunc(pattern, fn)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
