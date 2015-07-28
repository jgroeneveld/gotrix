package expenses

import (
	"net/http"

	"gotrix/app/model"
	"gotrix/lib/web/ctx"
	"gotrix/lib/web/form"
	"gotrix/lib/web/httperr"
	"gotrix/web/frontend/views"

	"gotrix/app/service"
	"gotrix/lib/errors"
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

	// TODO render shit
	return errors.New(" TODO render view for %v", expense)
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
