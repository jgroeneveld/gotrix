package db

import "database/sql"

type TxManagerFactory interface {
	Create() TxManager
}

func NewTxManagerFactory(con *sql.DB) TxManagerFactory {
	return &ConTxManagerFactory {
		con: con,
	}
}

type ConTxManagerFactory struct {
	con *sql.DB
}

func (f *ConTxManagerFactory) Create() TxManager {
	return NewTxManager(f.con)
}
