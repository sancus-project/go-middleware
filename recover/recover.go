package recover

import (
	"go.sancus.dev/web"
	"go.sancus.dev/web/middleware"
)

func Recover(h web.ErrorHandlerFunc) web.MiddlewareHandlerFunc {
	return middleware.Recover(h)
}
