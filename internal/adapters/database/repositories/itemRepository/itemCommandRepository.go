package itemRepository

import (
	"database/sql"
	"game-inventory-management/internal/domain/item"

	"github.com/google/uuid"
)

func Create(db *sql.DB, item item.Item) error {
	_, err := db.Exec(
		"INSERT INTO item (id, inventory_id, external_reference) VALUES ($1, $2, $3)",
		item.Id, item.InventoryId, item.ExternalReference,
	)
	if err != nil {
		return err
	}
	return nil
}

func UpdateItemInventoryId(db *sql.DB, itemId uuid.UUID, newInventoryId uuid.UUID) error {
	query := "UPDATE item SET inventory_id = $1 WHERE id = $2"
	_, err := db.Exec(query, newInventoryId, itemId)
	return err
}
