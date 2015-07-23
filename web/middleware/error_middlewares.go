package middleware

import (
	"github.com/jgroeneveld/bookie2/web/ctx"
	"github.com/jgroeneveld/bookie2/web/httperr"
	"github.com/jgroeneveld/bookie2/web/util"
	"net/http"
)

func RenderErrorsAsJSON() Middleware {
	return MiddlewareFunc(renderErrorsAsJSONFunc)
}

func renderErrorsAsJSONFunc(next HTTPHandle) HTTPHandle {
	return func(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		httpErr := httperr.Convert(next(rw, r, c))
		if httpErr != nil {
			rw.WriteHeader(httpErr.Status)
			_ = util.RenderJSON(rw, httpErr)
			c.Printf("error_response=%s", httpErr.Error())
		}
		return nil
	}
}

func RenderErrorsAsHTML() Middleware {
	return MiddlewareFunc(renderErrorsAsHTMLFunc)
}

func renderErrorsAsHTMLFunc(next HTTPHandle) HTTPHandle {
	return func(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		httpErr := httperr.Convert(next(rw, r, c))
		if httpErr != nil {
			rw.WriteHeader(httpErr.Status)
			// TODO render error as html page
			c.Printf("respond with error=%s", httpErr.Error())
		}
		return nil
	}
}
