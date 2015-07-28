package service

import (
	"gotrix/app/db"
	"gotrix/app/model"
	"gotrix/app/model/validate"
	"gotrix/lib/logger"
)

func CreateExpense(l logger.Logger, con db.Con, params CreateExpenseParams) (*model.Expense, error) {
	expense := model.NewExpense()

	expense.Description = params.Description
	expense.Amount = params.Amount

	err := validate.Expense(expense)
	if err != nil {
		return nil, err
	}

	err = db.InsertExpense(con, expense)
	if err != nil {
		return nil, err
	}

	return expense, nil
}

func ListExpenses(l logger.Logger, con db.Con) (list []*model.Expense, err error) {
	err = db.IterateExpenses(con, func(e *model.Expense) error {
		list = append(list, e)
		return nil
	})
	return list, err
}

type CreateExpenseParams struct {
	Description string
	Amount      int
}
