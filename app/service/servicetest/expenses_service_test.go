package servicetest

import (
	"gotrix/app/db"
	"gotrix/lib/logger"
	"testing"

	"gotrix/app/db/dbtest"
	"gotrix/app/db/dbtest/fabricate"

	"gotrix/app/service"

	"github.com/jgroeneveld/trial/assert"
)

func TestCreateExpense_Success(t *testing.T) {
	tx := dbtest.BeginTx()
	defer tx.Rollback()

	expense, err := service.CreateExpense(logger.Discard, tx, service.CreateExpenseParams{
		Description: "Peter Lustig",
		Amount:      42,
	})

	assert.MustBeNil(t, err)
	assert.Equal(t, "Peter Lustig", expense.Description)
	assert.Equal(t, 42, expense.Amount)

	cnt, err := db.CountExpenses(tx)
	assert.MustBeNil(t, err)
	assert.Equal(t, 1, cnt)
}

func TestCreateExpense_Validation(t *testing.T) {
	tx := dbtest.BeginTx()
	defer tx.Rollback()

	_, err := service.CreateExpense(logger.Discard, tx, service.CreateExpenseParams{
		Description: "",
		Amount:      0,
	})

	assert.MustBeEqual(t, "ValidationError: Description: must be present, Amount: must be greater than 0", err.Error())

	cnt, err := db.CountExpenses(tx)
	assert.MustBeNil(t, err)
	assert.Equal(t, 0, cnt)
}

func TestListExpenses(t *testing.T) {
	tx := dbtest.BeginTx()
	defer tx.Rollback()

	_, err := fabricate.Expense(tx).Description("Important Expense").Exec()
	assert.MustBeNil(t, err)

	expenses, err := service.ListExpenses(logger.Discard, tx)
	assert.MustBeNil(t, err)
	assert.MustBeEqual(t, 1, len(expenses))
	assert.Equal(t, "Important Expense", expenses[0].Description)
}
