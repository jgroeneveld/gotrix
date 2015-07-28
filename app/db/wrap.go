package db

import (
	"database/sql"
	"gotrix/app/apperrors"
	"gotrix/lib/errors"
)

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
