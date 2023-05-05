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

func AddBalanceToWallet(
	walletId uuid.UUID,
	value int64,
	test bool,
	db *sql.DB,
	log *zap.SugaredLogger,
) (*uuid.UUID, error) {
	id := uuid.New()

	walletEvent := domain.WalletEventStore{
		Id:       id,
		WalletId: walletId,
		Type:     domain.Credit,
		Value:    value,
		Test:     test,
	}

	walletEventRepository := walletRepositoryAdapter.NewWalletEventStoreCommandRepository()
	err := walletEventRepository.CreateEvent(walletEvent, db)

	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}

	walletQueryRepository := walletRepository.NewWalletQueryRepository()
	wallet, err := walletQueryRepository.Get(walletId, db)

	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}

	wallet.Value += value

	walletRepository := walletRepositoryAdapter.NewWalletCommandRepository()
	err = walletRepository.Update(wallet, db)

	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}

	return &id, nil
}
