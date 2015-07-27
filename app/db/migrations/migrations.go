package migrations

import (
	"github.com/jgroeneveld/gotrix/lib/db"
	"github.com/jgroeneveld/gotrix/lib/db/migrations"
	"github.com/jgroeneveld/gotrix/lib/logger"
)

func Exec(l logger.Logger, con db.Con) error {
	return migrations.Exec(l, con, Migrations)
}

var Migrations migrations.Migrations
