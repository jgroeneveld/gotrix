package handlers

import (
	"github.com/jgroeneveld/bookie2/app/service/expenses"
	"github.com/jgroeneveld/bookie2/lib/web/ctx"
	"github.com/jgroeneveld/bookie2/lib/web/form"
	"github.com/jgroeneveld/bookie2/lib/web/httperr"
	"net/http"
	"github.com/jgroeneveld/bookie2/lib/errors"
	"github.com/jgroeneveld/bookie2/app/db"
)

func CreateExpense(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
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
	return errors.New(" TODO RENDER JSON")
}

func ListExpenses(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
	// TODO render json
	return db.AllExpenses()
}
