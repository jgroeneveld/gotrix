package db

import "database/sql"

type TxFactory interface {
	BeginTx() (Tx, error)
}

func NewTxFactory(con *sql.DB) TxFactory {
	return &ConTxFactory{
		con: con,
	}
}

type ConTxFactory struct {
	con *sql.DB
}

func (f *ConTxFactory) BeginTx() (Tx, error) {
	return f.con.Begin()
}
