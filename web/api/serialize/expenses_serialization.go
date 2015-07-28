package serialize

import (
	"encoding/json"
	"gotrix/app/model"
	"gotrix/lib/errors"
	"io"
)

func Expense(w io.Writer, expense *model.Expense) error {
	rep := &expenseRep{
		Amount:      expense.Amount,
		Description: expense.Description,
	}

	return errors.Wrap(json.NewEncoder(w).Encode(rep))
}

func ExpensesList(w io.Writer, expenses []*model.Expense) error {
	rep := []*expenseRep{}

	for _, expense := range expenses {
		rep = append(rep, &expenseRep{
			Amount:      expense.Amount,
			Description: expense.Description,
		})
	}

	return errors.Wrap(json.NewEncoder(w).Encode(rep))
}

type expenseRep struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}
