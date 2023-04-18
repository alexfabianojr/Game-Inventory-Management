package walletBusiness

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/adapters/database/repositories/walletRepository"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func BalanceExchangesBetweenWallets(
	walletIdPayer uuid.UUID,
	walletIdPayee uuid.UUID,
	value int64,
	test bool,
	db *sql.DB,
	log *zap.SugaredLogger,
) (*uuid.UUID, error) {
	id := uuid.New()

	walletPayer, err := walletRepository.Get(walletIdPayer, db)

	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}

	walletPayee, err := walletRepository.Get(walletIdPayee, db)

	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}

	walletPayer.Value -= value

	if walletPayer.Value < 0 {
		return nil, errors.New("Payer doesn't have enough balance")
	}

	walletPayee.Value += value

	err = walletRepository.Update(walletPayer, db)

	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}

	err = walletRepository.Update(walletPayee, db)

	if err != nil {
		log.Error(err)
		return nil, errors.New(err.Error())
	}

	return &id, nil
}
