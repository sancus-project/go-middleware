package recover

import (
	"go.sancus.dev/web"
	"go.sancus.dev/web/middleware/recover"
)

func Recover(h web.ErrorHandlerFunc) web.MiddlewareHandlerFunc {
	return recover.Recover(h)
}
