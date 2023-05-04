package queryProjection

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

const (
	selectByInventoryId = `SELECT 
				i.id as inventory_id,
				i.wallet_id,
				w.value,
				i.external_reference
			FROM inventory i 
			INNER JOIN wallet w ON w.id = i.wallet_id 
			WHERE i.id = $1`
)

type InventoryWithWalletQueryProjection struct {
	InventoryId       uuid.UUID
	WalletId          uuid.UUID
	Value             int64
	ExternalReference uuid.UUID
}

func GetInventoryWithWalletByInventoryId(
	id uuid.UUID,
	db *sql.DB,
) (InventoryWithWalletQueryProjection, error) {
	var queryProjection InventoryWithWalletQueryProjection
	err := db.QueryRow(selectByInventoryId, id).
		Scan(
			&queryProjection.InventoryId,
			&queryProjection.WalletId,
			&queryProjection.Value,
			&queryProjection.ExternalReference,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return InventoryWithWalletQueryProjection{}, fmt.Errorf(
				"inventory and wallet with id inventory %s not found", id.String(),
			)
		}
		return InventoryWithWalletQueryProjection{}, err
	}
	return queryProjection, nil
}
