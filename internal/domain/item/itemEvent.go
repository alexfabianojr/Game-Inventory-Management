package item

import (
	"github.com/google/uuid"
)

type ItemEventStore struct {
	Id                  uuid.UUID
	OccurredOn          int64
	Type                string
	SenderInventoryId   *uuid.UUID
	ReceiverInventoryId *uuid.UUID
	ItemId              uuid.UUID
	Test                bool
}

const (
	New      string = "create"
	Delete   string = "delete"
	TradeIn  string = "trade_in"
	TradeOut string = "trade_out"
)
