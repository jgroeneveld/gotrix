package db

import (
	"database/sql"

	"github.com/jgroeneveld/gotrix/lib/errors"
)

func Connect(databaseURL string, applicationName string) (*sql.DB, error) {
	u := databaseURL + "&application_name=" + applicationName

	con, err := sql.Open("postgres", u)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	if err := con.Ping(); err != nil {
		return nil, errors.Wrap(err)
	}

	con.SetMaxIdleConns(2)
	con.SetMaxOpenConns(20)
	if _, err := con.Exec("SET statement_timeout TO '1800s'"); err != nil {
		return nil, err
	}

	return con, nil
}
