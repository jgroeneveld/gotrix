package db

import (
	"gotrix/app/model"
	"gotrix/lib/db"
)

func CountExpenses(con Con) (int, error) {
	cnt := 0
	row := con.QueryRow("SELECT count(1) FROM expenses")
	err := row.Scan(&cnt)
	return cnt, wrap(err)
}

func InsertExpense(con Con, expense *model.Expense) error {
	_, err := con.Exec(
		"INSERT INTO expenses (description, amount) VALUES ($1,$2)",
		expense.Description, expense.Amount,
	)
	return wrap(err)
}

func IterateExpenses(con Con, f func(*model.Expense) error) error {
	return iterateExpenses(con, f)
}

func iterateExpenses(con Con, f func(*model.Expense) error, opts ...func(*db.Opt)) error {
	rows, err := db.QueryWithOpts(con, "SELECT id, description, amount FROM expenses", opts...)
	if err != nil {
		return wrap(err)
	}

	for rows.Next() {
		expense := model.NewExpense()
		err := rows.Scan(&expense.ID, &expense.Description, &expense.Amount)
		if err != nil {
			return wrap(err)
		}
		err = f(expense)
		if err != nil {
			return wrap(err)
		}
	}

	return wrap(rows.Err())
}
