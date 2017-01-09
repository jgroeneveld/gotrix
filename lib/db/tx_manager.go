package db

import (
	"database/sql"
	"gotrix/lib/errors"
)

type TxManager interface {
	Begin() (Con, error)
	Close(success bool) error
}

func NewTxManager(con *sql.DB) *SimpleTxManager {
	return &SimpleTxManager{
		con: con,
	}
}

type SimpleTxManager struct {
	con        *sql.DB
	tx         Tx
	didBeginTx bool
}

func (m *SimpleTxManager) Begin() (Con, error) {
	tx, err := m.con.Begin()
	if err != nil {
		return nil, errors.Wrap(err)
	}

	m.tx = tx
	m.didBeginTx = true
	return tx, nil
}

func (m *SimpleTxManager) Close(success bool) error {
	if m.didBeginTx {
		if success {
			return errors.Wrap(m.tx.Commit())
		}
		return errors.Wrap(m.tx.Rollback())
	}

	m.tx = nil
	m.didBeginTx = false
	return nil
}
