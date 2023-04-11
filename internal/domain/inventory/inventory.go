package inventory

import "github.com/google/uuid"

type Inventory struct {
	Id                uuid.UUID
	WalletId          uuid.UUID
	ExternalReference uuid.UUID
	Size              int64
}
