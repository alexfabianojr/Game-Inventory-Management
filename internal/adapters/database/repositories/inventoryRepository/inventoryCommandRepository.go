package inventoryRepository

import (
	"database/sql"

	"github.com/google/uuid"
)

const (
	getCreateQuery = "INSERT INTO inventory(id, wallet_id, external_reference) VALUES($1, $2, $3)"
)

func Create(externalReference string, id uuid.UUID, db *sql.DB) error {
	stmt, err := db.Prepare(getCreateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id.String(), id.String(), externalReference)
	if err != nil {
		return err
	}
	return nil
}
