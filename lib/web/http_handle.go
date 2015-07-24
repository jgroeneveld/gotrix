package web

import (
	"net/http"

	"github.com/jgroeneveld/gotrix/lib/web/ctx"
)

type HTTPHandle func(http.ResponseWriter, *http.Request, *ctx.Context) error
