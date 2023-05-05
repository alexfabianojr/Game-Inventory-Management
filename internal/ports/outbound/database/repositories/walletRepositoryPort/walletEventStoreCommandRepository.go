package walletRepositoryPort

import (
	"database/sql"
	"game-inventory-management/internal/domain/wallet"
)

type WalletEventStoreCommandRepository interface {
	CreateEvent(walletEvent wallet.WalletEventStore, db *sql.DB) error
}
