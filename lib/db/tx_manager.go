package db

import "github.com/jgroeneveld/gotrix/lib/errors"

func NewTxManager(txFac TxFactory) *TxManager {
	return &TxManager{
		txFactory: txFac,
	}
}

type TxManager struct {
	txFactory  TxFactory
	tx         Tx
	didBeginTx bool
}

func (m *TxManager) Begin() (Con, error) {
	tx, err := m.txFactory.BeginTx()
	if err != nil {
		return nil, errors.Wrap(err)
	}

	m.tx = tx
	m.didBeginTx = true
	return tx, nil
}

func (m *TxManager) Close(success bool) error {
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
