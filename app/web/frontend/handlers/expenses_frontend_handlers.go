package expenses

import (
	"net/http"

	"github.com/go-errors/errors"
	"github.com/jgroeneveld/gotrix/app/model"
	"github.com/jgroeneveld/gotrix/app/service/expenses"
	"github.com/jgroeneveld/gotrix/app/web/frontend/views"
	"github.com/jgroeneveld/gotrix/lib/web/ctx"
	"github.com/jgroeneveld/gotrix/lib/web/form"
	"github.com/jgroeneveld/gotrix/lib/web/httperr"
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
	return errors.New(" TODO render view")
}

func ListExpenses(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
	view := &views.ExpensesList{
		Expenses: []*model.Expense{
			&model.Expense{Description: "Fahrrad", Amount: 109900},
			&model.Expense{Description: "iPhone", Amount: 14999},
		},
	}

	return views.RenderWithLayout(rw, view)
}
