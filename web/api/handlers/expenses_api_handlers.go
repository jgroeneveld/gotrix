package handlers

import (
	"net/http"

	"gotrix/app/service"
	"gotrix/lib/web/ctx"
	"gotrix/lib/web/form"
	"gotrix/lib/web/httperr"
	"gotrix/web/api/serialize"
)

func CreateExpense(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
	form, err := form.New(r)
	if err != nil {
		return httperr.BadRequest(err.Error())
	}

	params := service.CreateExpenseParams{
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

	expense, err := service.CreateExpense(c.Logger, con, params)
	if err != nil {
		return err
	}

	rw.WriteHeader(http.StatusCreated)
	return serialize.Expense(rw, expense)
}

func ListExpenses(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
	con, err := c.TxManager.Begin()
	if err != nil {
		return err
	}

	expenses, err := service.ListExpenses(c.Logger, con)
	if err != nil {
		return err
	}

	rw.WriteHeader(http.StatusOK)
	return serialize.ExpensesList(rw, expenses)
}
