package router

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/jgroeneveld/gotrix/lib/logger"
	"github.com/jgroeneveld/gotrix/lib/web"
	"github.com/jgroeneveld/gotrix/lib/web/ctx"
	"github.com/jgroeneveld/gotrix/lib/web/middleware"
	"github.com/julienschmidt/httprouter"
)

// chain handler with middleware
func httprouterAdapter(globalLogger logger.Logger) func(*middleware.Chain, web.HTTPHandle) httprouter.Handle {
	adapter := httpHandleConverter(globalLogger)
	return func(middlewares *middleware.Chain, handle web.HTTPHandle) httprouter.Handle {
		f := middlewares.Bind(handle)
		return adapter(f)
	}
}

func httpHandleConverter(globalLogger logger.Logger) func(web.HTTPHandle) httprouter.Handle {
	return func(handle web.HTTPHandle) httprouter.Handle {
		return func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
			// put request id into logger and headers to allow better error reporting / debugging
			requestID := newRequestID()
			l := globalLogger.Fork("request_id=" + requestID)
			rw.Header().Add("x-request-id", requestID)

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
