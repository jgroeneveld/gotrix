package validate

import (
	"github.com/jgroeneveld/gotrix/app/model"
	"strings"
)

func Expense(expense *model.Expense) error {
	validator := NewValidator()

	if strings.TrimSpace(expense.Description) == "" {
		validator.Add("Description", "must be present")
	}

	if expense.Amount <= 0 {
		validator.Add("Amount", "must be greater than 0")
	}

	return validator.Err()
}
