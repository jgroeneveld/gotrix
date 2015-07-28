package migrations

import (
	"gotrix/lib/db"
	"gotrix/lib/db/migrations"
	"gotrix/lib/logger"
)

func Exec(l logger.Logger, con db.Con) error {
	return migrations.Exec(l, con, Migrations)
}

var Migrations migrations.Migrations
