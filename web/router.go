package web

import (
	"github.com/go-errors/errors"
	"github.com/jgroeneveld/bookie2/logger"
	"github.com/jgroeneveld/bookie2/web/ctx"
	"github.com/jgroeneveld/bookie2/web/handlers/expenses"
	"github.com/jgroeneveld/bookie2/web/middleware"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewRouter(l *logger.Logger) http.Handler {
	globalMiddlewares := middleware.NewChain(
		middleware.RequestLogger(),
	)

	apiMiddlewares := middleware.NewChain(
		globalMiddlewares,
		middleware.RenderErrorsAsJSON(),
		// TODO middleware.RecoverPanics(),
	)

	frontendMiddlewares := middleware.NewChain(
		globalMiddlewares,
		middleware.RenderErrorsAsHTML(),
		// TODO middleware.RecoverPanics(),
	)

	r := httprouter.New()
	mw := middleware.Wrapper(l)

	r.GET("/", mw(frontendMiddlewares, expenses.ListHandler))
	r.POST("/", mw(apiMiddlewares, expenses.CreateHandler))

	return r
}

func testhandler(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
	c.Printf("handler called")
	return errors.New("alles kaput")
}

type LogMiddleware struct {
	Msg string
}

func (mw *LogMiddleware) Call(next middleware.HTTPHandle) middleware.HTTPHandle {
	return func(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		c.Printf("%s Before", mw.Msg)
		err := next(rw, r, c)
		c.Printf("%s After", mw.Msg)
		return err
	}
}
