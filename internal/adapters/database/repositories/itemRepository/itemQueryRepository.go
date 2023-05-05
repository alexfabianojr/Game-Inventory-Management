package itemRepository

import (
	"database/sql"
	"fmt"
	domain "game-inventory-management/internal/domain/item"
	port "game-inventory-management/internal/ports/outbound/database/repositories/itemRepositoryPort"

	"github.com/google/uuid"
)

type ItemQueryRepositoryPostgreSQL struct{}

func NewItemQueryRepository() port.ItemQueryRepository {
	return ItemQueryRepositoryPostgreSQL{}
}

const (
	getAllItemsByInventoryIdQuery = `SELECT 
									i.id,
									i.external_reference 
								FROM item i 
								WHERE i.inventory_id = $1`
	getQuery = "SELECT id, inventory_id, external_reference FROM item WHERE id = $1"
)

func (ItemQueryRepositoryPostgreSQL) Get(id uuid.UUID, db *sql.DB) (domain.Item, error) {
	var item domain.Item
	err := db.QueryRow(getQuery, id).
		Scan(&item.Id, &item.InventoryId, &item.ExternalReference)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Item{}, fmt.Errorf("wallet with ID %s not found", id.String())
		}
		return domain.Item{}, err
	}
	return item, nil
}

func (ItemQueryRepositoryPostgreSQL) GetAllItemsByInventoryId(inventoryId uuid.UUID, db *sql.DB) ([]domain.Item, error) {
	stmt, err := db.Prepare(getAllItemsByInventoryIdQuery)
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
