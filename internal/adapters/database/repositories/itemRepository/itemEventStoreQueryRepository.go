package itemRepository

import (
	"database/sql"
	"game-inventory-management/internal/domain/item"
	"game-inventory-management/internal/ports/database/repositories/itemRepositoryPort"

	"github.com/google/uuid"
)

type ItemEventStoreQueryRepositoryPostgreSQL struct{}

func NewItemEventStoreQueryRepository() itemRepositoryPort.ItemEventStoreQueryRepository {
	return ItemEventStoreQueryRepositoryPostgreSQL{}
}

const (
	getAllEventsByItemIdQuery = `SELECT 
									id,
									occurred_on,
									type,
									sender_inventory_id,
									receiver_inventory_id,
									item_id,
									test 
								FROM item_event_store 
								WHERE item_id = $1`
)

func (ItemEventStoreQueryRepositoryPostgreSQL) GetAllEventsByItemId(itemId uuid.UUID, db *sql.DB) ([]item.ItemEventStore, error) {
	stmt, err := db.Prepare(getAllEventsByItemIdQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(itemId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	itemEvents := []item.ItemEventStore{}
	for rows.Next() {
		itemEvent := item.ItemEventStore{}
		err = rows.Scan(
			&itemEvent.Id,
			&itemEvent.OccurredOn,
			&itemEvent.Type,
			&itemEvent.SenderInventoryId,
			&itemEvent.ReceiverInventoryId,
			&itemEvent.ItemId,
			&itemEvent.Test,
		)
		if err != nil {
			return nil, err
		}
		itemEvents = append(itemEvents, itemEvent)
	}

	return itemEvents, nil
}
