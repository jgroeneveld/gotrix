package fabricate

import (
	"gotrix/app/db"
	"gotrix/app/model"
	"math/rand"
)

func Expense(con db.Con) *ExpenseFabricator {
	expense := model.NewExpense()
	expense.Amount = int(rand.Int31())
	expense.Description = "A description"

	return &ExpenseFabricator{con: con, expense: expense}
}

type ExpenseFabricator struct {
	con     db.Con
	expense *model.Expense
}

func (f *ExpenseFabricator) Set(fn func(*model.Expense)) *ExpenseFabricator {
	fn(f.expense)
	return f
}

func (f *ExpenseFabricator) Description(v string) *ExpenseFabricator {
	f.expense.Description = v
	return f
}

func (f *ExpenseFabricator) Amount(v int) *ExpenseFabricator {
	f.expense.Amount = v
	return f
}

func (f *ExpenseFabricator) Exec() (*model.Expense, error) {
	return f.expense, db.InsertExpense(f.con, f.expense)
}
