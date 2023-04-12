package inventory

import (
	"database/sql"
	"game-inventory-management/internal/adapters/database/repositories/inventoryRepository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Create(externalReference string, db *sql.DB) (*uuid.UUID, error) {
	id := uuid.New()
	inventoryRepository.Create(externalReference, id, db)
	return &id, nil
}
