package dbtest

import (
	"database/sql"
	"gotrix/app/db/migrations"
	"gotrix/cfg"
	"gotrix/lib/db"
	"gotrix/lib/db/dbtest"
	"gotrix/lib/logger"
)

var cachedTestCon *sql.DB

func NewTestCon() *sql.DB {
	if cachedTestCon == nil {
		con, err := db.Connect(cfg.Config.DatabaseURL, cfg.Config.ApplicationName)
		if err != nil {
			panic(err)
		}
		cachedTestCon = con
	}
	return cachedTestCon
}

func BeginTx() *sql.Tx {
	tx, err := NewTestCon().Begin()
	if err != nil {
		panic(err)
	}
	err = migrations.Exec(logger.Discard, tx)
	if err != nil {
		panic(err)
	}
	return tx
}

func NewTxManagerFactory() *dbtest.TxManagerFactory {
	return dbtest.NewTxManagerFactory(BeginTx())
}
