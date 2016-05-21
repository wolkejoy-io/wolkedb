package wolkedb

import "database/sql"

type Engine struct  {
	db *sql.DB
	dialect string
}

func New(db *sql.DB, dialect string) *Engine {
	return &Engine{db:db, dialect:dialect}
}

func (self *Engine) BeginTx() (*Tx, error) {
	sqlTx, txBeginErr := self.db.Begin()
	if txBeginErr != nil {
		return nil, txBeginErr
	}
	self.db.Driver()
	tx := Tx{}
	tx.tx = sqlTx
	tx.dialect = self.dialect
	return &tx, nil
}

