package db

import (
	"database/sql"
	"github.com/jgroeneveld/bookie2/lib/errors"
)

func AllExpenses() error {
	return errors.Wrap(sql.ErrNoRows)
}
