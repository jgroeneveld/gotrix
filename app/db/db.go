package db

import (
	"database/sql"
	apperrors "github.com/jgroeneveld/gotrix/app/errors"
	"github.com/jgroeneveld/gotrix/lib/errors"
)

type Con interface {
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Prepare(query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// wrap wraps and converts database errors into application level errors
func wrap(err error) error {
	if err == nil {
		return nil
	}

	if err == sql.ErrNoRows {
		return apperrors.RecordNotFound()
	}

	return errors.Wrap(err)
}
