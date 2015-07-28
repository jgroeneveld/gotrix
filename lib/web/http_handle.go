package web

import (
	"net/http"

	"gotrix/lib/web/ctx"
)

type HTTPHandle func(http.ResponseWriter, *http.Request, *ctx.Context) error
