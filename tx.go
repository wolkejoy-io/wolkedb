package bitdb

import "database/sql"

type Tx struct  {
	tx *sql.Tx
	dialect string
}

func (self *Tx) Commit() error {
	return self.tx.Commit()
}

func (self *Tx) Rollback() error {
	return self.tx.Rollback()
}
