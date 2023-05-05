package projection

import (
	"database/sql"
	inventoryRepositoryAdapter "game-inventory-management/internal/adapters/database/repositories/inventoryRepository"
	"game-inventory-management/internal/adapters/database/repositories/itemRepository"
	"game-inventory-management/internal/domain/item"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetInventoryWithItems(
	inventoryId uuid.UUID,
	db *sql.DB,
	log *zap.SugaredLogger,
) (InventoryWithItemsProjection, error) {
	inventoryRepository := inventoryRepositoryAdapter.NewInventoryQueryRepositoryImpl()

	inventory, err := inventoryRepository.Get(inventoryId, db)
	if err != nil {
		log.Error(err)
		return InventoryWithItemsProjection{}, err
	}

	items, err := itemRepository.GetAllItemsByInventoryId(inventoryId, db)

	if err != nil {
		log.Error(err)
		return InventoryWithItemsProjection{}, err
	}

	projection := InventoryWithItemsProjection{
		Id:                         inventory.Id,
		InventoryWalletId:          inventory.WalletId,
		InventoryExternalReference: inventory.ExternalReference,
		Items:                      ParseToItemProjection(items),
	}

	return projection, nil
}

func ParseToItemProjection(items []item.Item) []ItemProjection {
	itemsProjection := []ItemProjection{}

	for _, item := range items {
		itemsProjection = append(
			itemsProjection,
			ItemProjection{
				Id:                item.Id,
				ExternalReference: item.ExternalReference,
			},
		)
	}
	return itemsProjection
}

type InventoryWithItemsProjection struct {
	Id                         uuid.UUID        `json:"id"`
	InventoryWalletId          uuid.UUID        `json:"inventory_wallet_id"`
	InventoryExternalReference uuid.UUID        `json:"inventory_external_reference"`
	Items                      []ItemProjection `json:"items"`
}

type ItemProjection struct {
	Id                uuid.UUID `json:"id"`
	ExternalReference uuid.UUID `json:"external_reference"`
}
