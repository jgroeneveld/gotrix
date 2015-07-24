package web

import (
	"github.com/go-errors/errors"
	"github.com/jgroeneveld/bookie2/lib/logger"
	apihandlers "github.com/jgroeneveld/bookie2/app/web/api/handlers"
	frontendhandlers "github.com/jgroeneveld/bookie2/app/web/frontend/handlers"
	"github.com/jgroeneveld/bookie2/lib/web/ctx"
	"github.com/jgroeneveld/bookie2/lib/web/middleware"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewRouter(l logger.Logger) http.Handler {
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

	r.GET("/expenses", mw(frontendMiddlewares, frontendhandlers.ListExpenses))
	r.POST("/expenses", mw(frontendMiddlewares, frontendhandlers.CreateExpense))

	r.GET("/api/v1/expenses", mw(apiMiddlewares, apihandlers.ListExpenses))
	r.POST("/api/v1/expenses", mw(apiMiddlewares, apihandlers.CreateExpense))

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
