package structable

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Transaction(DB *sql.DB, f func() error) error {
	var err error
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	err = f()
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

