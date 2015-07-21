package web

import (
	"github.com/jgroeneveld/bookie2/logger"
	"github.com/jgroeneveld/bookie2/web/ctx"
	"github.com/jgroeneveld/bookie2/web/handlers/expenses"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func NewRouter(l *logger.Logger) http.Handler {
	mw := ctx.Middleware(l)

	r := httprouter.New()

	r.GET("/", mw(expenses.ListHandler))
	r.POST("/", mw(expenses.CreateHandler))

	return r
}
