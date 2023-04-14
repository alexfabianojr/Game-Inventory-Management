package itemRepository

import (
	"database/sql"
	"game-inventory-management/internal/domain/item"
)

func CreateEvent(itemEvent item.ItemEventStore, db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO item_event_store "+
			"(id, occurred_on, type, third_party_inventory_id, wallet_event_store_id, external_reference, test) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7)",
		itemEvent.Id,
		itemEvent.OccurredOn,
		itemEvent.Type,
		itemEvent.ThirdPartyInventoryId,
		itemEvent.WalletEventStoreId,
		itemEvent.ExternalReference,
		itemEvent.Test,
	)
	if err != nil {
		return err
	}
	return nil
}
