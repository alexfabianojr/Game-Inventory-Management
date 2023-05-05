package walletRepositoryPort

import (
	"database/sql"
	"game-inventory-management/internal/domain/wallet"

	"github.com/google/uuid"
)

type WalletEventStoreQueryRepository interface {
	GetAllEventsByWalletId(walletId uuid.UUID, db *sql.DB) ([]wallet.WalletEventStore, error)
}
