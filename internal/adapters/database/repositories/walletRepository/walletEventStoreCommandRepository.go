package walletRepository

import (
	"database/sql"
	"game-inventory-management/internal/domain/wallet"
)

const (
	getCreateEventQuery = `INSERT INTO wallet_event_store 
								(id, wallet_id, type, third_party_wallet_id, value, test) 
							VALUES ($1, $2, $3, $4, $5, $6)`
)

func CreateEvent(walletEvent wallet.WalletEventStore, db *sql.DB) error {
	_, err := db.Exec(
		getCreateEventQuery,
		walletEvent.Id,
		walletEvent.WalletId,
		walletEvent.Type,
		walletEvent.ThirdPartyWalletId,
		walletEvent.Value,
		walletEvent.Test,
	)
	if err != nil {
		return err
	}
	return nil
}
