package middleware

import (
	"net/http"

	"github.com/mitchellh/go-server-timing"
)

// Middleware accessor for github.com/mitchellh/go-server-timing
// for compatibility with chi.Use()
func ServerTimer(next http.Handler) http.Handler {
	return servertiming.Middleware(next, nil)
}
