package walletRepository

import (
	"database/sql"
	"fmt"
	domain "game-inventory-management/internal/domain/wallet"

	"github.com/google/uuid"
)

func Get(id uuid.UUID, db *sql.DB) (domain.Wallet, error) {
	var wallet domain.Wallet
	err := db.QueryRow("SELECT id, value FROM wallet WHERE id = $1", id).Scan(&wallet.Id, &wallet.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Wallet{}, fmt.Errorf("wallet with ID %s not found", id.String())
		}
		return domain.Wallet{}, err
	}
	return wallet, nil
}
