package expenses

import (
	"github.com/jgroeneveld/bookie2/app/service/expenses"
	"github.com/jgroeneveld/bookie2/web/ctx"
	"github.com/jgroeneveld/bookie2/web/httperr"
	"github.com/jgroeneveld/bookie2/web/util/form"
	"net/http"
)

func CreateHandler(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
	form, err := form.New(r)
	if err != nil {
		return httperr.BadRequest(err.Error())
	}

	params := expenses.CreateParams{
		Description: form.ReqString("description"),
		Amount:      form.ReqInt("amount"),
	}

	if err := form.Err(); err != nil {
		return err
	}

	err = expenses.Create(c.Logger, params)
	if err != nil {
		return err
	}

	// TODO render shit
	return nil
}
