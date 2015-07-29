package expenses

import (
	"net/http"

	"gotrix/lib/web/ctx"
	"gotrix/lib/web/form"
	"gotrix/lib/web/httperr"
	"gotrix/web/frontend/views"

	"gotrix/app/apperrors"
	"gotrix/app/service"
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

	con, err := c.TxManager.Begin()
	if err != nil {
		return err
	}

	_, err = service.CreateExpense(c.Logger, con, params)
	if err != nil {
		ve, ok := apperrors.IsValidationError(err)
		if !ok {
			return err
		}
		view := &views.ExpenseForm{
			Description: params.Description,
			Amount:      params.Amount,
			Errors:      ve.FieldErrors,
		}
		return views.RenderWithLayout(rw, view)
	}

	http.Redirect(rw, r, "/expenses", 302)
	return nil
}

func ListExpenses(rw http.ResponseWriter, r *http.Request, c *ctx.Context) error {
	tx, err := c.TxManager.Begin()
	if err != nil {
		return err
	}
	expenses, err := service.ListExpenses(c.Logger, tx)
	if err != nil {
		return err
	}

	view := new(views.ExpensesList)
	for _, exp := range expenses {
		view.Expenses = append(view.Expenses, &views.ExpensesListItem{
			Description: exp.Description,
			Amount:      exp.Amount,
		})
	}

	return views.RenderWithLayout(rw, view)
}
