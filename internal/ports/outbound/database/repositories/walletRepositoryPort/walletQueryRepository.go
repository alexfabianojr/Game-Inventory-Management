package walletRepositoryPort

import (
	"database/sql"
	domain "game-inventory-management/internal/domain/wallet"

	"github.com/google/uuid"
)

type WalletQueryRepository interface {
	Get(id uuid.UUID, db *sql.DB) (domain.Wallet, error)
}
