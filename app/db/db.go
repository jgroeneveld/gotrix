package db

import (
	"database/sql"
	apperrors "gotrix/app/errors"
	"gotrix/lib/errors"
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
