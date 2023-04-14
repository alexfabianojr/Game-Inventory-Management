package itemBusiness

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/adapters/database/repositories/itemRepository"
	"game-inventory-management/internal/domain/item"
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

	itemEvent := item.ItemEventStore{
		Id:                id,
		OccurredOn:        time.Now().Unix(),
		Type:              item.New,
		ExternalReference: externalReference,
		Test:              false,
	}

	err := itemRepository.CreateEvent(itemEvent, db)

	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}

	item := item.Item{
		Id:                id,
		InventoryId:       inventoryId,
		ExternalReference: externalReference,
	}

	itemRepository.Create(db, item)

	return &id, nil
}
