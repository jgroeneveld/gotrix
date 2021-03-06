package middleware

import (
	"net/http"

	"gotrix/lib/errors"
	"gotrix/lib/web"
	"gotrix/lib/web/ctx"
)

func RecoverPanics() Middleware {
	return MiddlewareFunc(recoverPanicsFunc)
}

func recoverPanicsFunc(next web.HTTPHandle) web.HTTPHandle {
	return func(rw http.ResponseWriter, r *http.Request, c *ctx.Context) (err error) {
		defer func() {
			// TODO maybe we should skip the additional stack lines.
			if r := recover(); r != nil {
				err = errors.New("panic: %s", r)
			}
		}()
		return next(rw, r, c)
	}
}
