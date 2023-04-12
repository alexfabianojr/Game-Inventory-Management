package inventoryRepository

import (
	"database/sql"

	"github.com/google/uuid"
)

func Create(externalReference string, id uuid.UUID, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO inventory(id, wallet_id, external_reference) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, id, externalReference)
	if err != nil {
		return err
	}
	return nil
}
