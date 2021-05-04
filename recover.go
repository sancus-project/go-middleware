package middleware

import (
	"go.sancus.dev/web"
	"go.sancus.dev/web/middleware"
)

func Recoverer(h web.ErrorHandlerFunc) web.MiddlewareHandlerFunc {
	return middleware.Recoverer(h)
}
