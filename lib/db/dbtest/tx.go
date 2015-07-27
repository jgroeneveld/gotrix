package dbtest

import (
	"database/sql"

	"github.com/jgroeneveld/gotrix/app/db/migrations"
	"github.com/jgroeneveld/gotrix/lib/logger"
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
