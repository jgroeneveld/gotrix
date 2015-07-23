package middleware

import (
	"github.com/jgroeneveld/bookie2/logger"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/jgroeneveld/bookie2/web/shared/ctx"
	"crypto/rand"
	"fmt"
)

func Wrapper(l *logger.Logger) func(*Chain, HTTPHandle) httprouter.Handle {
	first := ForHTTPRouter(l)

	return func(middlewares *Chain, handle HTTPHandle) httprouter.Handle {
		f := middlewares.Bind(handle)
		return first(f)
	}
}

func ForHTTPRouter(globalLogger *logger.Logger) func(HTTPHandle) httprouter.Handle {
	return func(handle HTTPHandle) httprouter.Handle {
		return func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
			l := globalLogger.Fork("request_id=" + newRequestID())

			c := ctx.NewContext(l, params)
			err := handle(rw, r, c)

			if err != nil {
				panic("errors should not bubble up to this point")
			}
		}
	}
}

func newRequestID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)[0:8]
}
