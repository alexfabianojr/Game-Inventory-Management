package itemBusiness

import (
	"database/sql"
	itemRepositoryAdapter "game-inventory-management/internal/adapters/database/repositories/itemRepository"
	"game-inventory-management/internal/domain/item"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetItemEvents(
	itemId uuid.UUID,
	db *sql.DB,
	log *zap.SugaredLogger,
) ([]item.ItemEventStore, error) {
	itemRepository := itemRepositoryAdapter.NewItemEventStoreQueryRepository()
	events, err := itemRepository.GetAllEventsByItemId(itemId, db)
	if err != nil {
		log.Error(err)
		return []item.ItemEventStore{}, err
	}
	return events, nil
}
