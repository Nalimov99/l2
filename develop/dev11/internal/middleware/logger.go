package middleware

import (
	"context"
	"dev11/internal/platform/web"
	"log"
	"net/http"
)

func Logger(log *log.Logger) web.Middleware {
	// This is actual mw function to be executed
	f := func(before web.Handler) web.Handler {
		// This is main handler
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			err := before(ctx, w, r)

			v, ok := ctx.Value(web.KeyValues).(*web.ContexValues)
			if !ok {
				return web.ErrContextValueMissing
			}

			log.Printf(
				"%d %s %s",
				v.StatusCode, r.Method, r.URL.Path,
			)

			// Return the error to be handled further up the chain
			return err
		}

		return h
	}

	return f
}
