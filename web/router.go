package web

import (
	"net/http"

	"gotrix/lib/db"
	"gotrix/lib/logger"
	"gotrix/lib/web/middleware"
	"gotrix/lib/web/router"
	apihandlers "gotrix/web/api/handlers"
	"gotrix/web/frontend/assets"
	frontendhandlers "gotrix/web/frontend/handlers"
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
