package itemRepositoryPort

import (
	"database/sql"
	"game-inventory-management/internal/domain/item"
)

type ItemEventStoreCommandRepository interface {
	CreateEvent(itemEvent item.ItemEventStore, db *sql.DB) error
}
