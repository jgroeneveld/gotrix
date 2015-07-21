package views

import (
	"github.com/jgroeneveld/bookie2/app/model"
	"io"
)

type ExpensesList struct {
	Expenses []*model.Expense
}

func (v *ExpensesList) Render(w io.Writer) error {
	return writeExpensesList(w, v)
}
