package views

import "io"

type ExpensesListItem struct {
	Description string
	Amount      int
}

type ExpensesList struct {
	Expenses []*ExpensesListItem
}

func (v *ExpensesList) Render(w io.Writer) error {
	return writeExpensesList(w, v)
}
