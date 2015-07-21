package ctx

import (
	"crypto/rand"
	"fmt"
	"github.com/jgroeneveld/bookie2/logger"
	"github.com/jgroeneveld/bookie2/web/httperr"
	"github.com/jgroeneveld/bookie2/web/util"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func Middleware(globalLogger *logger.Logger) func(httpHandle) httprouter.Handle {
	return func(f httpHandle) httprouter.Handle {
		return func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
			l := globalLogger.Fork("request_id=" + newRequestID())

			c := NewContext(l, params)

			c.Printf("starting %s %s", r.Method, r.URL)
			startedAt := time.Now()

			httpErr := httperr.Convert(f(rw, r, c))
			if httpErr != nil {
				rw.WriteHeader(httpErr.Status)
				_ = util.RenderJSON(rw, httpErr)
			}

			suffix := ""
			if httpErr != nil {
				suffix = fmt.Sprintf(" ERROR=%s", httpErr.Error())
			}
			c.Printf("finished in %.06f%s", time.Since(startedAt).Seconds(), suffix)
		}
	}
}

type httpHandle func(http.ResponseWriter, *http.Request, *Context) error

func newRequestID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return fmt.Sprintf("%x", b)[0:8]
}
