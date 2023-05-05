package walletBusiness

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/adapters/database/repositories/walletRepository"
	walletRepositoryAdapter "game-inventory-management/internal/adapters/database/repositories/walletRepository"
	"game-inventory-management/internal/domain/wallet"

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
) error {
	walletPayer, err := walletRepository.Get(walletIdPayer, db)

	if err != nil {
		log.Error(err)
		return errors.New(err.Error())
	}

	walletPayee, err := walletRepository.Get(walletIdPayee, db)

	if err != nil {
		log.Error(err)
		return errors.New(err.Error())
	}

	walletPayer.Value -= value

	if walletPayer.Value < 0 {
		return errors.New("Payer doesn't have enough balance")
	}

	eventPayer := wallet.WalletEventStore{
		Id:                 uuid.New(),
		WalletId:           walletIdPayer,
		Type:               wallet.TradeOut,
		ThirdPartyWalletId: &walletIdPayee,
		Value:              value,
		Test:               test,
	}

	eventPayee := wallet.WalletEventStore{
		Id:                 uuid.New(),
		WalletId:           walletIdPayee,
		Type:               wallet.TradeIn,
		ThirdPartyWalletId: &walletIdPayer,
		Value:              value,
		Test:               test,
	}

	walletEventRepository := walletRepositoryAdapter.NewWalletEventStoreCommandRepository()
	walletEventRepository.CreateEvent(eventPayer, db)
	walletEventRepository.CreateEvent(eventPayee, db)

	walletPayee.Value += value

	walletRepository := walletRepositoryAdapter.NewWalletCommandRepository()
	err = walletRepository.Update(walletPayer, db)

	if err != nil {
		log.Error(err)
		return errors.New(err.Error())
	}

	err = walletRepository.Update(walletPayee, db)

	if err != nil {
		log.Error(err)
		return errors.New(err.Error())
	}

	return nil
}
