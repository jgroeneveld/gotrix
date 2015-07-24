package web

import (
	"net/http"

	"github.com/go-errors/errors"
	apihandlers "github.com/jgroeneveld/gotrix/app/web/api/handlers"
	frontendhandlers "github.com/jgroeneveld/gotrix/app/web/frontend/handlers"
	"github.com/jgroeneveld/gotrix/lib/logger"
	"github.com/jgroeneveld/gotrix/lib/web"
	"github.com/jgroeneveld/gotrix/lib/web/ctx"
	"github.com/jgroeneveld/gotrix/lib/web/middleware"
	"github.com/jgroeneveld/gotrix/lib/web/router"
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

	r := router.New(l)

	r.Get("/expenses", frontendMiddlewares, frontendhandlers.ListExpenses)
	r.Post("/expenses", frontendMiddlewares, frontendhandlers.CreateExpense)

	r.Get("/api/v1/expenses", apiMiddlewares, apihandlers.ListExpenses)
	r.Post("/api/v1/expenses", apiMiddlewares, apihandlers.CreateExpense)

	return r
}

func testhandler(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
	c.Printf("handler called")
	return errors.New("alles kaput")
}

type LogMiddleware struct {
	Msg string
}

func (mw *LogMiddleware) Call(next web.HTTPHandle) web.HTTPHandle {
	return func(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		c.Printf("%s Before", mw.Msg)
		err := next(rw, r, c)
		c.Printf("%s After", mw.Msg)
		return err
	}
}
