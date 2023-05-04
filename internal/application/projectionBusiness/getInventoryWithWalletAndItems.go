package projection

import (
	"database/sql"
	"game-inventory-management/internal/adapters/database/repositories/itemRepository"
	projectionRepository "game-inventory-management/internal/adapters/database/repositories/queryProjectionRepository"
	"game-inventory-management/internal/domain/item"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetInventoryWithWalletAndItems(
	inventoryId uuid.UUID,
	db *sql.DB,
	log *zap.SugaredLogger,
) (InventoryWithWalletAndItemsProjection, error) {
	inventoryWithWallet, err := projectionRepository.GetInventoryWithWalletByInventoryId(inventoryId, db)
	if err != nil {
		log.Error(err)
		return InventoryWithWalletAndItemsProjection{}, err
	}

	items, err := itemRepository.GetAllItemsByInventoryId(inventoryWithWallet.InventoryId, db)
	if err != nil {
		log.Error(err)
		return InventoryWithWalletAndItemsProjection{}, err
	}

	return ParseToInventoryWithWalletAndItemProjection(inventoryWithWallet, items), nil
}

type InventoryWithWalletAndItemsProjection struct {
	InventoryId       uuid.UUID        `json:"inventory_id"`
	WalletId          uuid.UUID        `json:"wallet_id"`
	Value             int64            `json:"value"`
	ExternalReference uuid.UUID        `json:"external_reference"`
	Items             []ItemProjection `json:"items"`
}

func ParseToInventoryWithWalletAndItemProjection(
	projection projectionRepository.InventoryWithWalletQueryProjection,
	items []item.Item,
) InventoryWithWalletAndItemsProjection {
	return InventoryWithWalletAndItemsProjection{
		InventoryId:       projection.InventoryId,
		WalletId:          projection.WalletId,
		Value:             projection.Value,
		ExternalReference: projection.ExternalReference,
		Items:             ParseToItemProjection(items),
	}
}
