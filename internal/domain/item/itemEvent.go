package item

import (
	"github.com/google/uuid"
)

type ItemEventStore struct {
	Id                    uuid.UUID
	OccurredOn            int64
	Type                  string
	ThirdPartyInventoryId *uuid.UUID
	WalletEventStoreId    *uuid.UUID
	ExternalReference     uuid.UUID
	Test                  bool
}

const (
	New    string = "create"
	Delete string = "delete"
	Trade  string = "trade"
)
