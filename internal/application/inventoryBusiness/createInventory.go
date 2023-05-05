package inventoryBusiness

import (
	"database/sql"
	"errors"
	"fmt"
	"game-inventory-management/internal/adapters/database/repositories/inventoryRepository"
	walletRepositoryAdapter "game-inventory-management/internal/adapters/database/repositories/walletRepository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func Create(externalReference string, db *sql.DB, log *zap.SugaredLogger) (*uuid.UUID, error) {
	id := uuid.New()
	log.Info(fmt.Sprintf("Creating new inventory and wallet for %s", id))

	walletRepository := walletRepositoryAdapter.NewWalletCommandRepository()
	err := walletRepository.Create(id, db)
	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}
	log.Info(fmt.Sprintf("Created wallet for id %s", id))

	repository := inventoryRepository.NewInventoryCommandRepository()
	err = repository.Create(externalReference, id, db)
	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}
	log.Info(fmt.Sprintf("Created inventory for id %s", id))

	return &id, nil
}
