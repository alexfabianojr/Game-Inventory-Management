package item

import (
	"github.com/google/uuid"
)

type ItemEvent struct {
	Id                    uuid.UUID
	OccurredOn            int64
	Type                  string
	ThirdPartyInventoryId *uuid.UUID
	TradeReference        *uuid.UUID
	Value                 *int64
	Test                  bool
}

const (
	New    string = "create"
	Delete string = "delete"
	Trade  string = "trade"
)
