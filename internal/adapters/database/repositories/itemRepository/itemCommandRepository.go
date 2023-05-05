package itemRepository

import (
	"database/sql"
	"game-inventory-management/internal/domain/item"
	"game-inventory-management/internal/ports/outbound/database/repositories/itemRepositoryPort"

	"github.com/google/uuid"
)

type ItemCommandRepositoryPostgreSQL struct{}

func NewItemCommandRepositoryImpl() itemRepositoryPort.ItemCommandRepository {
	return ItemCommandRepositoryPostgreSQL{}
}

const (
	getCreateQuery                = "INSERT INTO item (id, inventory_id, external_reference) VALUES ($1, $2, $3)"
	getUpdateItemInventoryIdQuery = "UPDATE item SET inventory_id = $1 WHERE id = $2"
)

func (ItemCommandRepositoryPostgreSQL) Create(db *sql.DB, item item.Item) error {
	_, err := db.Exec(
		getCreateQuery,
		item.Id, item.InventoryId, item.ExternalReference,
	)
	if err != nil {
		return err
	}
	return nil
}

func (ItemCommandRepositoryPostgreSQL) UpdateItemInventoryId(db *sql.DB, itemId uuid.UUID, newInventoryId uuid.UUID) error {
	_, err := db.Exec(getUpdateItemInventoryIdQuery, newInventoryId, itemId)
	return err
}
