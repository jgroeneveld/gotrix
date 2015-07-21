package model

type Expense struct {
	Description string
	Amount      int
}

func NewExpense() *Expense {
	return &Expense{}
}
