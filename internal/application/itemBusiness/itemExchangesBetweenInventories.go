package itemBusiness

import (
	"database/sql"
	"errors"
	inventoryRepositoryAdapter "game-inventory-management/internal/adapters/database/repositories/inventoryRepository"
	itemRepositoryAdapter "game-inventory-management/internal/adapters/database/repositories/itemRepository"
	domain "game-inventory-management/internal/domain/item"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func ItemExchangesBetweenInventories(
	itemExchanges ItemExchanges,
	db *sql.DB,
	log *zap.SugaredLogger,
) error {
	inventoryRepository := inventoryRepositoryAdapter.NewInventoryQueryRepository()

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

	itemRepository := itemRepositoryAdapter.NewItemQueryRepository()
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

	itemEventRepository := itemRepositoryAdapter.NewItemEventStoreCommandRepository()
	err = itemEventRepository.CreateEvent(tradeOutEvent, db)

	if err != nil {
		log.Error(err)
		return errors.New(err.Error())
	}

	tradeInEvent := domain.ItemEventStore{
		Id:                  uuid.New(),
		OccurredOn:          time.Now().Unix(),
		Type:                domain.TradeIn,
		SenderInventoryId:   &itemExchanges.SenderInventoryId,
		ReceiverInventoryId: &itemExchanges.ReceiverInventoryId,
		ItemId:              itemExchanges.ItemId,
		Test:                itemExchanges.Test,
	}

	err = itemEventRepository.CreateEvent(tradeInEvent, db)

	if err != nil {
		log.Error(err)
		return errors.New(err.Error())
	}

	itemCommandRepository := itemRepositoryAdapter.NewItemCommandRepositoryImpl()
	err = itemCommandRepository.UpdateItemInventoryId(db, itemExchanges.ItemId, itemExchanges.ReceiverInventoryId)

	if err != nil {
		log.Error(err)
		return errors.New(err.Error())
	}

	return nil
}

type ItemExchanges struct {
	SenderInventoryId   uuid.UUID `json:"sender_inventory_id"`
	ReceiverInventoryId uuid.UUID `json:"receiver_inventory_id"`
	ItemId              uuid.UUID `json:"item_id"`
	Test                bool      `json:"test"`
}
