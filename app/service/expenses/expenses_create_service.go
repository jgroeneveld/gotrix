package expenses

import (
	"gotrix/app/db"
	"gotrix/app/model"
	"gotrix/app/model/validate"
	"gotrix/lib/logger"
)

type CreateParams struct {
	Description string
	Amount      int
}

func Create(l logger.Logger, con db.Con, params CreateParams) error {
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
