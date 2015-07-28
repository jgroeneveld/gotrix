package dbtest

import (
	"database/sql"

	"gotrix/app/db/migrations"
	"gotrix/lib/logger"
)

func MustBeginTx(con *sql.DB) *sql.Tx {
	tx, err := con.Begin()
	if err != nil {
		panic(err)
	}

	err = migrations.Exec(logger.Discard, tx)
	if err != nil {
		panic(err)
	}

	return tx
}
