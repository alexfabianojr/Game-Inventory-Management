package commandEndpoints

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func StartCommandRouter(e *echo.Echo, db *sql.DB, log *zap.SugaredLogger) {
	e.POST("/command/inventory/create/:externalReference", createInventoryHandler(db, log))

	e.POST("/command/item/create/:externalReference/:inventoryId", createItemHandler(db, log))
	e.POST("/command/item/exchanges", itemExchangesBetweenInventories(db, log))

	e.POST("/command/wallet/add-balance/:walletId", addBalanceToWalletHandler(db, log))
	e.POST("/command/wallet/balance-exchanges", balanceExchangesBetweenWallets(db, log))
	// reload wallet balance
}
