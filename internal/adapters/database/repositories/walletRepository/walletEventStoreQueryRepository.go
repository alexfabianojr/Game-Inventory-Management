package walletRepository

import (
	"database/sql"
	"game-inventory-management/internal/domain/wallet"

	"github.com/google/uuid"
)

type WalletEventStoreQueryRepositoryPostgreSQL struct{}

const (
	getAllEventsByWalletIdQuery = "SELECT id, wallet_id, type, third_party_wallet_id, value, test FROM wallet_event_store WHERE wallet_id = $1"
)

func GetAllEventsByWalletId(walletId uuid.UUID, db *sql.DB) ([]wallet.WalletEventStore, error) {
	stmt, err := db.Prepare(getAllEventsByWalletIdQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(walletId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	walletEvents := []wallet.WalletEventStore{}
	for rows.Next() {
		walletEvent := wallet.WalletEventStore{}
		err = rows.Scan(
			&walletEvent.Id,
			&walletEvent.WalletId,
			&walletEvent.Type,
			&walletEvent.ThirdPartyWalletId,
			&walletEvent.Value,
			&walletEvent.Test,
		)
		if err != nil {
			return nil, err
		}
		walletEvents = append(walletEvents, walletEvent)
	}

	return walletEvents, nil
}
