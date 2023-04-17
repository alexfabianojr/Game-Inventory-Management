package itemRepository

import (
	"database/sql"
	"game-inventory-management/internal/domain/item"
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
