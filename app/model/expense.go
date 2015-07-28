package model

type Expense struct {
	ID          int
	Description string
	Amount      int
}

func NewExpense() *Expense {
	return &Expense{}
}
