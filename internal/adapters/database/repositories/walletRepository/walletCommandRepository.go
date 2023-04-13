package walletRepository

import (
	"database/sql"

	"github.com/google/uuid"
)

func Create(id uuid.UUID, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO wallet(id, value) VALUES($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id.String(), 0)
	if err != nil {
		return err
	}
	return nil
}
