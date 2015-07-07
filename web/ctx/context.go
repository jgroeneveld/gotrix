package ctx

import (
	"github.com/jgroeneveld/bookie2/logger"
	"github.com/julienschmidt/httprouter"
)

type Context struct {
	Logger *logger.Logger
}

func (c *Context) Printf(format string, args ...interface{}) {
	c.Logger.Printf(format, args...)
}

func NewContext(l *logger.Logger, params httprouter.Params) *Context {
	c := &Context{
		Logger: l,
	}
	return c
}
