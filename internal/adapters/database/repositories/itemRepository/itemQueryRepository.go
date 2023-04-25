package itemRepository

import (
	"database/sql"
	"fmt"
	domain "game-inventory-management/internal/domain/item"

	"github.com/google/uuid"
)

func Get(id uuid.UUID, db *sql.DB) (domain.Item, error) {
	var item domain.Item
	err := db.QueryRow("SELECT * FROM item WHERE id = $1", id).Scan(&item)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Item{}, fmt.Errorf("wallet with ID %s not found", id.String())
		}
		return domain.Item{}, err
	}
	return item, nil
}
