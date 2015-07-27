package db

import (
	"database/sql"

	"github.com/jgroeneveld/gotrix/lib/errors"
)

func WithTx(con *sql.DB, f func(Tx) error) error {
	tx, err := con.Begin()
	if err != nil {
		return errors.Wrap(err)
	}

	err = f(tx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	cerr := tx.Commit()
	return errors.Wrap(cerr)
}
