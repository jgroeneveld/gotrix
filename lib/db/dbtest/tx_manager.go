package dbtest

import (
	"database/sql"

	"gotrix/lib/db"
	"gotrix/lib/errors"
)

func NewTxManager(tx *sql.Tx) *TxManager {
	return &TxManager{
		Tx: tx,
	}
}

type TxManager struct {
	Tx                 *sql.Tx
	CloseSuccessCalled bool
	CloseFailCalled    bool
	txOpened bool
	txClosed bool
}

func (f *TxManager) Begin() (db.Con, error) {
	if f.txOpened {
		return nil, errors.New("tx already opened")
	}
	f.txOpened = true
	return f.Tx, nil
}

func (m *TxManager) Close(success bool) error {
	switch {
	case !m.txOpened:
		return errors.New("no tx opened")
	case m.txClosed:
		return errors.New("tx already closed")
	case success:
		m.CloseSuccessCalled = true
	default:
		m.CloseFailCalled = true
	}
	m.txClosed = true
	return nil
}

func (m *TxManager) Rollback() {
	err := m.Tx.Rollback()
	if err != nil {
		panic(err)
	}
}
