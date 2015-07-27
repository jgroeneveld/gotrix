package main

import (
	"github.com/dynport/dgtk/log"
	"github.com/jgroeneveld/gotrix/app/db/migrations"
	"github.com/jgroeneveld/gotrix/cfg"
	"github.com/jgroeneveld/gotrix/lib/db"
	"github.com/jgroeneveld/gotrix/lib/errors"
	"github.com/jgroeneveld/gotrix/lib/logger"
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
