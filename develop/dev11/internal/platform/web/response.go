package web

import (
	"context"
	"encoding/json"
	"net/http"
)

type CommonRespond[T any] struct {
	Result T `json:"result"`
}

// Respond marshals value to a JSON and send it to the client
func Respond(ctx context.Context, w http.ResponseWriter, val interface{}, statusCode int) error {
	v, ok := ctx.Value(KeyValues).(*ContexValues)
	if !ok {
		return ErrContextValueMissing
	}
	v.StatusCode = statusCode

	if val == nil {
		w.WriteHeader(statusCode)
		return nil
	}

	data, err := json.Marshal(val)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	if _, err := w.Write(data); err != nil {
		return err
	}

	return nil
}

// Respond error knows how to handle errors going out to the client
func RespondError(ctx context.Context, w http.ResponseWriter, err error) error {
	// If the error was of the type *Error the handles
	// has a specific status code and error to run
	if webErr, ok := err.(*Error); ok {
		resp := ErrorResponse{
			Error: webErr.Err.Error(),
		}

		return Respond(ctx, w, resp, webErr.Status)
	}

	resp := ErrorResponse{
		Error: http.StatusText(http.StatusServiceUnavailable),
	}

	return Respond(ctx, w, resp, http.StatusServiceUnavailable)
}
