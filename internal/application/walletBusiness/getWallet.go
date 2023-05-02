package walletBusiness

import (
	"database/sql"
	"game-inventory-management/internal/adapters/database/repositories/walletRepository"
	domain "game-inventory-management/internal/domain/wallet"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetWallet(
	id uuid.UUID,
	db *sql.DB,
	log *zap.SugaredLogger,
) (domain.Wallet, error) {
	wallet, err := walletRepository.Get(id, db)
	if err != nil {
		log.Error(err)
		return domain.Wallet{}, err
	}
	return wallet, nil
}
