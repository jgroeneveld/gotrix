package dbtest

import (
	"database/sql"

	"gotrix/lib/db"
)

func NewTxManager(tx *sql.Tx) *TxManager {
	return &TxManager{
		Tx: tx,
	}
}

type TxManager struct {
	Tx                 *sql.Tx
	closeSuccessCalled bool
	closeFailCalled    bool
}

func (f *TxManager) Begin() (db.Con, error) {
	return f.Tx, nil
}

func (m *TxManager) Close(success bool) error {
	if success {
		m.closeSuccessCalled = true
	} else {
		m.closeFailCalled = true
	}
	return nil
}

func (m *TxManager) Rollback() {
	err := m.Tx.Rollback()
	if err != nil {
		panic(err)
	}
}
