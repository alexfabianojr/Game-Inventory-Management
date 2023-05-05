package inventoryRepository

import (
	"database/sql"
	"fmt"
	domain "game-inventory-management/internal/domain/inventory"
	port "game-inventory-management/internal/ports/database/repositories/inventoryRepositoryPort"

	"github.com/google/uuid"
)

type InventoryQueryRepositoryImpl struct{}

func NewInventoryQueryRepositoryImpl() port.InventoryQueryRepository {
	return InventoryQueryRepositoryImpl{}
}

const (
	getQuery                    = "SELECT id, external_reference, wallet_id FROM inventory WHERE id = $1"
	getByExternalReferenceQuery = "SELECT id, external_reference, wallet_id FROM inventory WHERE external_reference = $1"
)

func (InventoryQueryRepositoryImpl) Get(id uuid.UUID, db *sql.DB) (domain.Inventory, error) {
	var inventory domain.Inventory
	err := db.QueryRow(getQuery, id).
		Scan(&inventory.Id, &inventory.ExternalReference, &inventory.WalletId)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Inventory{}, fmt.Errorf("Inventory with ID %s not found", id.String())
		}
		return domain.Inventory{}, err
	}

	return inventory, nil
}

func (InventoryQueryRepositoryImpl) GetByExternalReference(externalReference uuid.UUID, db *sql.DB) (domain.Inventory, error) {
	var inventory domain.Inventory
	err := db.QueryRow(getByExternalReferenceQuery, externalReference).
		Scan(&inventory.Id, &inventory.ExternalReference, &inventory.WalletId)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Inventory{}, fmt.Errorf("Inventory with ID %s not found", externalReference.String())
		}
		return domain.Inventory{}, err
	}

	return inventory, nil
}
