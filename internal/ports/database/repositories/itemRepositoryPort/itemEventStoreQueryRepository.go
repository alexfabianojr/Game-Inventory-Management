package itemRepositoryPort

import (
	"database/sql"
	"game-inventory-management/internal/domain/item"

	"github.com/google/uuid"
)

type ItemEventStoreQueryRepository interface {
	GetAllEventsByItemId(itemId uuid.UUID, db *sql.DB) ([]item.ItemEventStore, error)
}
