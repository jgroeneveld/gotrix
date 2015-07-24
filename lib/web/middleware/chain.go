package middleware

import "github.com/jgroeneveld/gotrix/lib/web"

type Middleware interface {
	Bind(next web.HTTPHandle) web.HTTPHandle
}

type MiddlewareFunc func(next web.HTTPHandle) web.HTTPHandle

func (f MiddlewareFunc) Bind(next web.HTTPHandle) web.HTTPHandle {
	return f(next)
}

func NewChain(middlewares ...Middleware) *Chain {
	return &Chain{Middlewares: middlewares}
}

type Chain struct {
	Middlewares []Middleware
}

func (chain *Chain) Bind(next web.HTTPHandle) web.HTTPHandle {
	for i := len(chain.Middlewares) - 1; i >= 0; i-- {
		next = chain.Middlewares[i].Bind(next)
	}
	return next
}
