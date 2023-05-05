package walletRepository

import (
	"database/sql"
	domain "game-inventory-management/internal/domain/wallet"
	port "game-inventory-management/internal/ports/outbound/database/repositories/walletRepositoryPort"

	"github.com/google/uuid"
)

type WalletCommandRepositoryPostgreSQL struct{}

func NewWalletCommandRepository() port.WalletCommandRepository {
	return WalletCommandRepositoryPostgreSQL{}
}

const (
	getCreateQuery = "INSERT INTO wallet(id, value) VALUES($1, $2)"
	getUpdateQuery = "UPDATE wallet SET value = $1 WHERE id = $2"
)

func (WalletCommandRepositoryPostgreSQL) Create(id uuid.UUID, db *sql.DB) error {
	stmt, err := db.Prepare(getCreateQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id.String(), 0)
	if err != nil {
		return err
	}
	return nil
}

func (WalletCommandRepositoryPostgreSQL) Update(wallet domain.Wallet, db *sql.DB) error {
	_, err := db.Exec(getUpdateQuery, wallet.Value, wallet.Id)
	if err != nil {
		return err
	}
	return nil
}
