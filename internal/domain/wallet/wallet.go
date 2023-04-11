package wallet

import "github.com/google/uuid"

type Wallet struct {
	Id                uuid.UUID
	InventoryId       uuid.UUID
	LastWalletEventId uuid.UUID
	Value             int64
}
