package projection

import (
	"database/sql"
	projectionRepository "game-inventory-management/internal/adapters/database/repositories/queryProjectionRepository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetInventoryWithWalletByInventoryId(
	id uuid.UUID,
	db *sql.DB,
	log *zap.SugaredLogger,
) (InventoryWithWalletProjection, error) {
	projection, err := projectionRepository.GetInventoryWithWalletByInventoryId(id, db)
	if err != nil {
		log.Error(err)
		return InventoryWithWalletProjection{}, err
	}
	result := ParseProjection(projection)
	return result, nil
}

type InventoryWithWalletProjection struct {
	InventoryId       uuid.UUID `json:"inventory_id"`
	WalletId          uuid.UUID `json:"wallet_id"`
	Value             int64     `json:"value"`
	ExternalReference uuid.UUID `json:"external_reference"`
}

func ParseProjection(projection projectionRepository.InventoryWithWalletQueryProjection) InventoryWithWalletProjection {
	return InventoryWithWalletProjection{
		InventoryId:       projection.InventoryId,
		WalletId:          projection.WalletId,
		Value:             projection.Value,
		ExternalReference: projection.ExternalReference,
	}
}
