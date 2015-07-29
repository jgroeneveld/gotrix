package views

import "io"

type ExpenseForm struct {
	Description string
	Amount      int
	Errors      map[string][]string
}

func (v *ExpenseForm) Render(w io.Writer) error {
	return writeExpenseForm(w, v)
}
