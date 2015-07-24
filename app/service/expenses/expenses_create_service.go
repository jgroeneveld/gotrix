package expenses

import (
	"github.com/jgroeneveld/gotrix/app/model"
	"github.com/jgroeneveld/gotrix/app/model/validate"
	"github.com/jgroeneveld/gotrix/lib/logger"
)

type CreateParams struct {
	Description string
	Amount      int
}

func Create(l logger.Logger, params CreateParams) error {
	expense := model.NewExpense()

	expense.Description = params.Description
	expense.Amount = params.Amount

	err := validate.Expense(expense)
	if err != nil {
		return err
	}

	// TODO do some real work
	return nil
}
