package ctx

import (
	"crypto/rand"
	"fmt"
	"github.com/jgroeneveld/bookie2/logger"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

func Middleware(l *logger.Logger) func(httpHandle) httprouter.Handle {
	return func(f httpHandle) httprouter.Handle {
		return func(rw http.ResponseWriter, r *http.Request, params httprouter.Params) {
			l = l.Fork("request_id=" + newRequestID())

			c := NewContext(l, params)

			c.Printf("starting %s %s", r.Method, r.URL)
			startedAt := time.Now()

			err := f(rw, r, c)

			suffix := ""
			if err != nil {
				suffix = fmt.Sprintf(" ERROR=%s", err.Error())
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
