package middleware

import (
	"encoding/json"
	"github.com/jgroeneveld/bookie2/lib/web/ctx"
	"github.com/jgroeneveld/bookie2/lib/web/httperr"
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
			_ = renderJSON(rw, httpErr)
			c.Printf("error_response=%s", httpErr.Error())
			if stack := httpErr.Stacktrace; stack != "" {
				c.Printf("%s", stack)
			}
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

func renderJSON(w http.ResponseWriter, i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(b)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte("\n"))

	return err
}
