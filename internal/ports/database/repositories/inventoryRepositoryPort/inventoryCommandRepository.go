package inventoryRepositoryPort

import (
	"database/sql"

	"github.com/google/uuid"
)

type InventoryCommandRepository interface {
	Create(externalReference string, id uuid.UUID, db *sql.DB) error
}
