package itemBusiness

import (
	"database/sql"
	"errors"
	itemRepositoryAdapter "game-inventory-management/internal/adapters/database/repositories/itemRepository"
	domain "game-inventory-management/internal/domain/item"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func Create(
	inventoryId uuid.UUID,
	externalReference uuid.UUID,
	test bool,
	db *sql.DB,
	log *zap.SugaredLogger,
) (*uuid.UUID, error) {
	id := uuid.New()

	item := domain.Item{
		Id:                id,
		InventoryId:       inventoryId,
		ExternalReference: externalReference,
	}

	itemRepository := itemRepositoryAdapter.NewItemCommandRepositoryImpl()
	err := itemRepository.Create(db, item)

	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}

	itemEvent := domain.ItemEventStore{
		Id:         id,
		OccurredOn: time.Now().Unix(),
		Type:       domain.New,
		ItemId:     item.Id,
		Test:       test,
	}

	itemEventRepository := itemRepositoryAdapter.NewItemEventStoreCommandRepository()
	err = itemEventRepository.CreateEvent(itemEvent, db)

	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}

	return &id, nil
}
