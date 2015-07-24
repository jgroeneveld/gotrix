package db

import (
	"database/sql"
	"github.com/jgroeneveld/gotrix/app/errors"
)

func AllExpenses() error {
	return errors.RecordNotFound()
}

func randomerror() error {
	return wrap(sql.ErrNoRows)
}
