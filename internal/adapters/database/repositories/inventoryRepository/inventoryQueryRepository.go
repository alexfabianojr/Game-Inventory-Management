package inventoryRepository

import (
	"database/sql"
	"fmt"
	domain "game-inventory-management/internal/domain/inventory"

	"github.com/google/uuid"
)

func Get(id uuid.UUID, db *sql.DB) (domain.Inventory, error) {
	var inventory domain.Inventory
	err := db.QueryRow("SELECT * FROM item WHERE id = $1", id).Scan(&inventory)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Inventory{}, fmt.Errorf("Inventory with ID %s not found", id.String())
		}
		return domain.Inventory{}, err
	}
	return inventory, nil
}