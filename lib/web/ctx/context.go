package ctx

import (
	"gotrix/lib/db"
	"gotrix/lib/logger"
	"github.com/julienschmidt/httprouter"
)

type Context struct {
	Logger    logger.Logger
	Params    httprouter.Params
	TxManager db.TxManager
}

func (c *Context) Printf(format string, args ...interface{}) {
	c.Logger.Printf(format, args...)
}

func (c *Context) ParamByName(name string) string {
	return c.Params.ByName(name)
}

func NewContext(l logger.Logger, params httprouter.Params) *Context {
	c := &Context{
		Logger: l,
		Params: params,
	}
	return c
}
