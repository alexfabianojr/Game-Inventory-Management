package wallet

import (
	"github.com/google/uuid"
)

type WalletEventStore struct {
	Id                 uuid.UUID
	WalletId           uuid.UUID
	Type               string
	ThirdPartyWalletId *uuid.UUID
	Value              int64
	Test               bool
}

const (
	Credit   string = "credit"
	Debit    string = "debit"
	TradeIn  string = "trade_in"
	TradeOut string = "trade_out"
)
