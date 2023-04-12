package slot

import (
	"github.com/google/uuid"
)

type Slot struct {
	Id          uuid.UUID
	ItemId      uuid.UUID
	InventoryId uuid.UUID
}
