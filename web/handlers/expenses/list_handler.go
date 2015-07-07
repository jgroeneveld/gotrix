package expenses

import (
	"github.com/jgroeneveld/bookie2/web/ctx"
	"io"
	"net/http"
	"strings"
)

func ListHandler(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
	io.Copy(rw, strings.NewReader("<h1>Hello World</h1>"))
	return nil
}
