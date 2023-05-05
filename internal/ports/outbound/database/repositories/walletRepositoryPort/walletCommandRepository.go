package walletRepositoryPort

import (
	"database/sql"
	domain "game-inventory-management/internal/domain/wallet"

	"github.com/google/uuid"
)

type WalletCommandRepository interface {
	Create(id uuid.UUID, db *sql.DB) error
	Update(wallet domain.Wallet, db *sql.DB) error
}
