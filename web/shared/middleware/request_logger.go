package middleware

import (
	"github.com/jgroeneveld/bookie2/web/shared/ctx"
	"net/http"
	"time"
)

func RequestLogger() Middleware {
	return MiddlewareFunc(requestLoggerFunc)
}

func requestLoggerFunc(next HTTPHandle) HTTPHandle {
	return func(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		c.Printf("starting %s %s", r.Method, r.URL)
		startedAt := time.Now()
		err := next(rw, r, c)
		c.Printf("finished in %.06f", time.Since(startedAt).Seconds())
		return err
	}
}
