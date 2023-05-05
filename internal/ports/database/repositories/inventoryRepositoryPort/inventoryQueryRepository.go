package inventoryRepositoryPort

import (
	"database/sql"
	domain "game-inventory-management/internal/domain/inventory"

	"github.com/google/uuid"
)

type InventoryQueryRepository interface {
	Get(id uuid.UUID, db *sql.DB) (domain.Inventory, error)
	GetByExternalReference(externalReference uuid.UUID, db *sql.DB) (domain.Inventory, error)
}
