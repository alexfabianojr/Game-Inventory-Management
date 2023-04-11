package item

import (
	"github.com/google/uuid"
)

type Item struct {
	id                uuid.UUID
	inventoryId       uuid.UUID
	externalReference uuid.UUID
}
