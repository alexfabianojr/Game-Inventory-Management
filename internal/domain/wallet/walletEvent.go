package wallet

import "github.com/google/uuid"

type WalletEvent struct {
	Id       uuid.UUID
	WalletId uuid.UUID
	Type     WalletEventType
	Trade    *TradeWalletEvent
	Value    int64
	Test     bool
}

type WalletEventType string

const (
	Credit WalletEventType = "credit"
	Debit  WalletEventType = "debit"
	Trade  WalletEventType = "trade"
)

type TradeWalletEvent struct {
	FromWalletId uuid.UUID
	ToWalletId   uuid.UUID
	Value        int64
}
