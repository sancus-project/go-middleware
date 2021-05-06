package middleware

import (
	"net/http"

	"github.com/mitchellh/go-server-timing"

	"go.sancus.dev/web"
)

// Middleware accessor for github.com/mitchellh/go-server-timing
// for compatibility with chi.Use()
func ServerTimer(next http.Handler) http.Handler {
	return servertiming.Middleware(next, nil)
}

// Middleware to produce a Server-Timer metric
func ServerTimerMetric(name string) web.MiddlewareHandlerFunc {
	if len(name) == 0 {
		return func(next http.Handler) http.Handler {
			return next
		}
	}

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if t := servertiming.FromContext(r.Context(); t != nil {
				defer t.NewMetric(name).Start().Stop()
			}
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
