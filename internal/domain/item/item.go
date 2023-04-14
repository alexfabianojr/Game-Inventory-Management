package item

import (
	"github.com/google/uuid"
)

type Item struct {
	Id                uuid.UUID
	InventoryId       uuid.UUID
	ExternalReference uuid.UUID
}
