package bitdb

import "database/sql"

type BitDB struct  {
	db *sql.DB
	dialect string
}

func New(db *sql.DB, dialect string) *BitDB {
	return &BitDB{db:db, dialect:dialect}
}

func (self *BitDB) BeginTx() (*Tx, error) {
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

