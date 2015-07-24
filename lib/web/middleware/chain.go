package middleware

import (
	"github.com/jgroeneveld/gotrix/lib/web/ctx"
	"net/http"
)

type HTTPHandle func(http.ResponseWriter, *http.Request, *ctx.Context) error

type Middleware interface {
	Bind(next HTTPHandle) HTTPHandle
}

type MiddlewareFunc func(next HTTPHandle) HTTPHandle

func (f MiddlewareFunc) Bind(next HTTPHandle) HTTPHandle {
	return f(next)
}

func NewChain(middlewares ...Middleware) *Chain {
	return &Chain{Middlewares: middlewares}
}

type Chain struct {
	Middlewares []Middleware
}

func (chain *Chain) Bind(next HTTPHandle) HTTPHandle {
	for i := len(chain.Middlewares) - 1; i >= 0; i-- {
		next = chain.Middlewares[i].Bind(next)
	}
	return next
}
