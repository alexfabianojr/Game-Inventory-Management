package itemRepository

import (
	"database/sql"
	"game-inventory-management/internal/domain/item"
)

func CreateEvent(itemEvent item.ItemEventStore, db *sql.DB) error {
	_, err := db.Exec(
		"INSERT INTO item_event_store "+
			"(id, occurred_on, type, sender_inventory_id, "+
			"receiver_inventory_id, item_id, test) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		itemEvent.Id,
		itemEvent.OccurredOn,
		itemEvent.Type,
		itemEvent.SenderInventoryId,
		itemEvent.ReceiverInventoryId,
		itemEvent.ItemId,
		itemEvent.Test,
	)
	if err != nil {
		return err
	}
	return nil
}
