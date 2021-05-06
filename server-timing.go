package middleware

import (
	"net/http"

	"github.com/mitchellh/go-server-timing"

	"go.sancus.dev/web"
	"go.sancus.dev/web/middleware"
)

// Middleware to enable Server-Time if a given function agrees,
// or unconditionally if none given
func ServerTimer(allowed func(r *http.Request) bool) web.MiddlewareHandlerFunc {

	timed := func(next http.Handler) http.Handler {
		return servertiming.Middleware(next, nil)
	}

	if allowed == nil {
		// unconditional
		return timed
	}

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if allowed(r) {
				next = timed(next)
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

// Middleware to produce a Server-Timer metric
func ServerTimerMetric(name string) web.MiddlewareHandlerFunc {
	if len(name) == 0 {
		return middleware.NOP
	}

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if t := servertiming.FromContext(r.Context()); t != nil {
				defer t.NewMetric(name).Start().Stop()
			}
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
