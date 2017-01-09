package middleware

import (
	"net/http"

	"gotrix/lib/db"
	"gotrix/lib/web"
	"gotrix/lib/web/ctx"
)

// TxMiddleware injects a TxManager into the Context and manages the transaction based onrequest result.
// If the request was a success, it tries to commit the transaction and reports any commit errors
// otherwise it rollsback and logs any rollback errors.
func TxMiddleware(txMFac db.TxManagerFactory) Middleware {
	return &txMiddleware{
		TxManagerFactory: txMFac,
	}
}

type txMiddleware struct {
	TxManagerFactory db.TxManagerFactory
}

func (mw *txMiddleware) Bind(next web.HTTPHandle) web.HTTPHandle {
	return func(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
		c.TxManager = mw.TxManagerFactory.Create()

		err := next(rw, r, c)

		if err != nil {
			cerr := c.TxManager.Close(false)
			if cerr != nil {
				c.Printf("error while closing tx %s", cerr.Error())
			}
			return err
		}

		cerr := c.TxManager.Close(true)
		if cerr != nil {
			return cerr
		}
		return nil
	}
}
