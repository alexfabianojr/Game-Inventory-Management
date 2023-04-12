package wallet

import (
	"github.com/google/uuid"
)

type Wallet struct {
	Id    uuid.UUID
	Value int64
}
