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
	err := db.InsertExpense(tx, expense)
	assert.MustBeNil(t, err)
}
