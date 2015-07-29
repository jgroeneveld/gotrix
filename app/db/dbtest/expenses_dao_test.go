package dbtest

import (
	"gotrix/app/db"
	"gotrix/app/model"
	"testing"

	"github.com/jgroeneveld/trial/assert"
)

func TestInsertExpense(t *testing.T) {
	tx := BeginTx()
	defer tx.Rollback()

	expense := model.NewExpense()
	expense.Description = "some description"
	expense.Amount = 23
	err := db.InsertExpense(tx, expense)
	assert.MustBeNil(t, err)

	var amount int
	var description string
	row := tx.QueryRow("SELECT description, amount FROM expenses")
	err = row.Scan(&description, &amount)

	assert.MustBeNil(t, err)
	assert.MustBeEqual(t, "some description", description)
	assert.MustBeEqual(t, 23, amount)
}
