package web

import (
	"net/http"

	"github.com/jgroeneveld/gotrix/lib/db"
	"github.com/jgroeneveld/gotrix/lib/logger"
	"github.com/jgroeneveld/gotrix/lib/web/middleware"
	"github.com/jgroeneveld/gotrix/lib/web/router"
	apihandlers "github.com/jgroeneveld/gotrix/web/api/handlers"
	"github.com/jgroeneveld/gotrix/web/frontend/assets"
	frontendhandlers "github.com/jgroeneveld/gotrix/web/frontend/handlers"
)

func NewRouter(l logger.Logger, txManager db.TxManager) http.Handler {
	afterErrorHandlingChain := middleware.NewChain(
		middleware.RequestLogger(),
	)

	beforeErrorHandlingChain := middleware.NewChain(
		middleware.TxMiddleware(txManager),
		middleware.RecoverPanics(),
	)

	apiMiddlewares := middleware.NewChain(
		afterErrorHandlingChain,
		middleware.RenderErrorsAsJSON(),
		beforeErrorHandlingChain,
	)

	frontendMiddlewares := middleware.NewChain(
		afterErrorHandlingChain,
		middleware.RenderErrorsAsHTML(),
		beforeErrorHandlingChain,
	)

	r := router.New(l)

	r.Get("/expenses", frontendMiddlewares, frontendhandlers.ListExpenses)
	r.Post("/expenses", frontendMiddlewares, frontendhandlers.CreateExpense)

	r.Get("/api/v1/expenses", apiMiddlewares, apihandlers.ListExpenses)
	r.Post("/api/v1/expenses", apiMiddlewares, apihandlers.CreateExpense)

	r.ServeFiles("/assets/*filepath", assets.FileSystem())

	return r
}
