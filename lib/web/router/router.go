package router

import (
	"net/http"

	"github.com/jgroeneveld/gotrix/lib/logger"
	"github.com/jgroeneveld/gotrix/lib/web"
	"github.com/jgroeneveld/gotrix/lib/web/middleware"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	router  *httprouter.Router
	adapter func(*middleware.Chain, web.HTTPHandle) httprouter.Handle
}

func New(l logger.Logger) *Router {
	r := new(Router)
	r.router = httprouter.New()
	r.adapter = HTTPRouterAdapter(l)
	return r
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(rw, req)
}

func (r *Router) Get(route string, mwc *middleware.Chain, handler web.HTTPHandle) {
	r.router.GET(route, r.adapter(mwc, handler))
}

func (r *Router) Post(route string, mwc *middleware.Chain, handler web.HTTPHandle) {
	r.router.POST(route, r.adapter(mwc, handler))
}

func (r *Router) Put(route string, mwc *middleware.Chain, handler web.HTTPHandle) {
	r.router.PUT(route, r.adapter(mwc, handler))
}

func (r *Router) Patch(route string, mwc *middleware.Chain, handler web.HTTPHandle) {
	r.router.PATCH(route, r.adapter(mwc, handler))
}

func (r *Router) Delete(route string, mwc *middleware.Chain, handler web.HTTPHandle) {
	r.router.DELETE(route, r.adapter(mwc, handler))
}

func (r *Router) Head(route string, mwc *middleware.Chain, handler web.HTTPHandle) {
	r.router.HEAD(route, r.adapter(mwc, handler))
}
