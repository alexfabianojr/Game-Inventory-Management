package itemRepository

import (
	"database/sql"
	"fmt"
	domain "game-inventory-management/internal/domain/item"

	"github.com/google/uuid"
)

const (
	getAllItemsByInventoryId = `SELECT 
									i.id,
									i.external_reference 
								FROM item i 
								WHERE i.inventory_id = $1`
)

func Get(id uuid.UUID, db *sql.DB) (domain.Item, error) {
	var item domain.Item
	err := db.QueryRow("SELECT id, inventory_id, external_reference FROM item WHERE id = $1", id).
		Scan(&item.Id, &item.InventoryId, &item.ExternalReference)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Item{}, fmt.Errorf("wallet with ID %s not found", id.String())
		}
		return domain.Item{}, err
	}
	return item, nil
}

func GetAllItemsByInventoryId(inventoryId uuid.UUID, db *sql.DB) ([]domain.Item, error) {
	stmt, err := db.Prepare(getAllItemsByInventoryId)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(inventoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []domain.Item{}
	for rows.Next() {
		item := domain.Item{}
		err = rows.Scan(&item.Id, &item.ExternalReference)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
