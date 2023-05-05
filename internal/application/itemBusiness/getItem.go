package itemBusiness

import (
	"database/sql"
	itemRepositoryAdapter "game-inventory-management/internal/adapters/database/repositories/itemRepository"
	domain "game-inventory-management/internal/domain/item"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetItem(
	id uuid.UUID,
	db *sql.DB,
	log *zap.SugaredLogger,
) (domain.Item, error) {
	itemRepository := itemRepositoryAdapter.NewItemQueryRepository()
	item, err := itemRepository.Get(id, db)
	if err != nil {
		log.Error(err)
		return domain.Item{}, err
	}
	return item, nil
}
