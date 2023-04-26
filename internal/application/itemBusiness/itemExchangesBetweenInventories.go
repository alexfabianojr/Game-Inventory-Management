package itemBusiness

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/adapters/database/repositories/inventoryRepository"
	"game-inventory-management/internal/adapters/database/repositories/itemRepository"
	domain "game-inventory-management/internal/domain/item"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func itemExchangesBetweenInventories(
	itemExchanges ItemExchangesBetweenInventories,
	db *sql.DB,
	log *zap.SugaredLogger,
) error {
	err := validateInventoryExists(itemExchanges, db, log)

	if err != nil {
		log.Error(err)
		return errors.New(err.Error())
	}

	item, err := itemRepository.Get(itemExchanges.ItemId, db)

	if err != nil {
		log.Error(err)
		return errors.New(err.Error())
	}

	if item.InventoryId != itemExchanges.SenderInventoryId {
		message := "The sender inventory doesn't owns the item"
		log.Error(message)
		return errors.New(message)
	}

	item.InventoryId = itemExchanges.ReceiverInventoryId

	tradeOutEvent := domain.ItemEventStore{
		Id:                  uuid.New(),
		OccurredOn:          time.Now().Unix(),
		Type:                domain.TradeOut,
		SenderInventoryId:   &itemExchanges.SenderInventoryId,
		ReceiverInventoryId: &itemExchanges.ReceiverInventoryId,
		ItemId:              itemExchanges.ItemId,
		Test:                itemExchanges.Test,
	}

	itemRepository.CreateEvent(tradeOutEvent, db)

	tradeInEvent := domain.ItemEventStore{
		Id:                  uuid.New(),
		OccurredOn:          time.Now().Unix(),
		Type:                domain.TradeIn,
		SenderInventoryId:   &itemExchanges.SenderInventoryId,
		ReceiverInventoryId: &itemExchanges.ReceiverInventoryId,
		ItemId:              itemExchanges.ItemId,
		Test:                itemExchanges.Test,
	}

	itemRepository.CreateEvent(tradeInEvent, db)

	itemRepository.UpdateItemInventoryId(db, itemExchanges.ItemId, itemExchanges.ReceiverInventoryId)

	return nil
}

func validateInventoryExists(
	itemExchanges ItemExchangesBetweenInventories,
	db *sql.DB,
	log *zap.SugaredLogger,
) error {
	_, err := inventoryRepository.Get(itemExchanges.ReceiverInventoryId, db)

	if err != nil {
		log.Error(err)
		return errors.New("Error while trying to find receiver inventory")
	}

	_, err = inventoryRepository.Get(itemExchanges.SenderInventoryId, db)

	if err != nil {
		log.Error(err)
		return errors.New("Error while trying to find sender inventory")
	}

	return nil
}

type ItemExchangesBetweenInventories struct {
	SenderInventoryId   uuid.UUID
	ReceiverInventoryId uuid.UUID
	ItemId              uuid.UUID
	Test                bool
}
