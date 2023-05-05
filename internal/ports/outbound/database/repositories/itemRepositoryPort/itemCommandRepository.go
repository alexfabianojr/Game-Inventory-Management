package itemRepositoryPort

import (
	"database/sql"
	"game-inventory-management/internal/domain/item"

	"github.com/google/uuid"
)

type ItemCommandRepository interface {
	Create(db *sql.DB, item item.Item) error
	UpdateItemInventoryId(db *sql.DB, itemId uuid.UUID, newInventoryId uuid.UUID) error
}
