package projection

import (
	"database/sql"
	"game-inventory-management/internal/adapters/database/repositories/inventoryRepository"
	"game-inventory-management/internal/adapters/database/repositories/itemRepository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetInventoryWithItems(
	id uuid.UUID,
	db *sql.DB,
	log *zap.SugaredLogger,
) (InventoryWithItemsProjection, error) {
	inventory, err := inventoryRepository.Get(id, db)
	if err != nil {
		log.Error(err)
		return InventoryWithItemsProjection{}, err
	}

	items, err := itemRepository.GetAllItemsByInventoryId(id, db)

	if err != nil {
		log.Error(err)
		return InventoryWithItemsProjection{}, err
	}

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

	projection := InventoryWithItemsProjection{
		Id:                         inventory.Id,
		InventoryWalletId:          inventory.WalletId,
		InventoryExternalReference: inventory.ExternalReference,
		Items:                      itemsProjection,
	}

	return projection, nil
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
