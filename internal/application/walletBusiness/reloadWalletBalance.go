package walletBusiness

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/adapters/database/repositories/walletRepository"
	walletRepositoryAdapter "game-inventory-management/internal/adapters/database/repositories/walletRepository"
	domain "game-inventory-management/internal/domain/wallet"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func ReloadWalletBalance(
	walletId uuid.UUID,
	test bool,
	db *sql.DB,
	log *zap.SugaredLogger,
) error {
	events, err := walletRepository.GetAllEventsByWalletId(walletId, db)

	if err != nil {
		log.Error(err)
		return errors.New(err.Error())
	}

	balance := CalculateWalletBalance(events)

	wallet := domain.Wallet{
		Id:    walletId,
		Value: balance,
	}

	walletRepository := walletRepositoryAdapter.NewWalletCommandRepository()
	err = walletRepository.Update(wallet, db)

	if err != nil {
		log.Error(err)
		return errors.New(err.Error())
	}

	return nil
}
