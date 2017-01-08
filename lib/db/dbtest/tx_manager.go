package dbtest

import (
	"database/sql"

	"gotrix/lib/db"
	"gotrix/lib/errors"
)

func NewTxManagerFactory(tx *sql.Tx) *TxManagerFactory {
	return &TxManagerFactory{
		Tx: tx,
	}
}

type TxManagerFactory struct {
	Tx *sql.Tx
}

func (txMFac *TxManagerFactory) Create() db.TxManager {
	return &TxManager{Tx: txMFac.Tx}
}

func (txMFac *TxManagerFactory) Close() error {
	return txMFac.Tx.Rollback()
}

type TxManager struct {
	Tx                 *sql.Tx
	CloseSuccessCalled bool
	CloseFailCalled    bool
	txOpened           bool
	txClosed           bool
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
