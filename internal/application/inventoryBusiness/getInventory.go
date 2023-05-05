package inventoryBusiness

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/adapters/database/repositories/inventoryRepository"
	domain "game-inventory-management/internal/domain/inventory"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetInventory(
	externalReference uuid.UUID,
	db *sql.DB,
	log *zap.SugaredLogger,
) (domain.Inventory, error) {
	repository := inventoryRepository.NewInventoryQueryRepositoryImpl()
	inventory, err := repository.GetByExternalReference(externalReference, db)
	if err != nil {
		log.Error(err)
		return domain.Inventory{}, errors.New(err.Error())
	}
	return inventory, nil
}
