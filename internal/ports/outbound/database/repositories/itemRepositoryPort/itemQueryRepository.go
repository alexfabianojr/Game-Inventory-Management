package itemRepositoryPort

import (
	"database/sql"
	domain "game-inventory-management/internal/domain/item"

	"github.com/google/uuid"
)

type ItemQueryRepository interface {
	Get(id uuid.UUID, db *sql.DB) (domain.Item, error)
	GetAllItemsByInventoryId(inventoryId uuid.UUID, db *sql.DB) ([]domain.Item, error)
}
