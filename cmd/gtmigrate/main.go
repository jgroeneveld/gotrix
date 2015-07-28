package main

import (
	"gotrix/app/db/migrations"
	"gotrix/cfg"
	"gotrix/lib/db"
	"gotrix/lib/errors"
	"gotrix/lib/logger"
	"log"
)

func main() {
	con, err := db.Connect(cfg.Config.DatabaseURL, cfg.Config.ApplicationName+"_gtmigrate")
	if err != nil {
		log.Fatal(errors.ErrorWithStack(err))
	}

	err = db.WithTx(con, func(tx db.Tx) error {
		l := logger.New(">>")
		return migrations.Exec(l, tx)
	})

	if err != nil {
		log.Fatal(errors.ErrorWithStack(err))
	}
}
