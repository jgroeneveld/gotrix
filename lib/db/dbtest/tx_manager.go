package dbtest

import (
	"database/sql"

	"github.com/jgroeneveld/gotrix/lib/db"
)

func NewTestTxManager(con *sql.DB) *TestTxManager {
	return &TestTxManager{
		Tx: MustBeginTx(con),
	}
}

type TestTxManager struct {
	Tx                 *sql.Tx
	closeSuccessCalled bool
	closeFailCalled    bool
}

func (f *TestTxManager) Begin() (db.Con, error) {
	return f.Tx, nil
}

func (m *TestTxManager) Close(success bool) error {
	if success {
		m.closeSuccessCalled = true
	} else {
		m.closeFailCalled = true
	}
	return nil
}

func (m *TestTxManager) Rollback() {
	err := m.Tx.Rollback()
	if err != nil {
		panic(err)
	}
}
