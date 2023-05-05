package walletBusiness

import (
	"database/sql"
	"game-inventory-management/internal/adapters/database/repositories/walletRepository"
	"game-inventory-management/internal/domain/wallet"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetWalletEvents(
	walletId uuid.UUID,
	db *sql.DB,
	log *zap.SugaredLogger,
) ([]wallet.WalletEventStore, error) {
	walletEventsRepository := walletRepository.NewWalletEventStoreQueryRepository()
	events, err := walletEventsRepository.GetAllEventsByWalletId(walletId, db)
	if err != nil {
		log.Error(err)
		return []wallet.WalletEventStore{}, err
	}
	return events, nil
}
