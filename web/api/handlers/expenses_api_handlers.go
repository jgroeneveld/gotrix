package handlers

import (
	"net/http"

	"gotrix/app/service/expenses"
	"gotrix/lib/errors"
	"gotrix/lib/web/ctx"
	"gotrix/lib/web/form"
	"gotrix/lib/web/httperr"
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

	con, err := c.TxManager.Begin()
	if err != nil {
		return err
	}

	err = expenses.Create(c.Logger, con, params)
	if err != nil {
		return err
	}

	// TODO render shit
	return errors.New(" TODO RENDER JSON")
}

func ListExpenses(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
	// TODO render json
	rw.Write([]byte("TODO content please"))
	return nil
}
