package middleware

import (
	"net/http"
	"time"

	"gotrix/lib/web"
	"gotrix/lib/web/ctx"
)

func RequestLogger() Middleware {
	return MiddlewareFunc(requestLoggerFunc)
}

func requestLoggerFunc(next web.HTTPHandle) web.HTTPHandle {
	return func(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		c.Printf("starting %s %s", r.Method, r.URL)
		startedAt := time.Now()
		err := next(rw, r, c)
		c.Printf("finished in %.06f", time.Since(startedAt).Seconds())
		return err
	}
}
