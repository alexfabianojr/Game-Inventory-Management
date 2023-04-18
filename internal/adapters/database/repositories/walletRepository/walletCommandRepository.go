package walletRepository

import (
	"database/sql"
	domain "game-inventory-management/internal/domain/wallet"

	"github.com/google/uuid"
)

func Create(id uuid.UUID, db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO wallet(id, value) VALUES($1, $2)")
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

func Update(wallet domain.Wallet, db *sql.DB) error {
	_, err := db.Exec("UPDATE wallet SET value = $1 WHERE id = $2", wallet.Value, wallet.Id)
	if err != nil {
		return err
	}
	return nil
}
