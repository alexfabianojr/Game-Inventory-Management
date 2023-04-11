package item

import "github.com/google/uuid"

type ItemEvent struct {
	Id         uuid.UUID
	OccurredOn int64
	Type       ItemEventType
	Trade      *TradeItemEvent
	Test       bool
}

type ItemEventType string

const (
	New    ItemEventType = "new"
	Delete ItemEventType = "delete"
	Trade  ItemEventType = "trade"
)

type TradeItemEvent struct {
	FromInventoryId uuid.UUID
	ToInventoryId   uuid.UUID
	Value           int64
}
