package queryEndpoint

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"go.uber.org/zap"
)

func StartQueryRouter(e *echo.Echo, db *sql.DB, log *zap.SugaredLogger) {
	e.GET("/query/inventory/:externalReference", getInventory(db, log))
	e.GET("/query/inventory/summary/wallet/:id", GetInventoryWithWallet(db, log))
	e.GET("/query/inventory/summary/wallet/items/:id", GetInventoryWithWalletAndItems(db, log))
	e.GET("/query/inventory/summary/items/:id", GetInventoryWithItems(db, log))

	e.GET("/query/wallet/:id", getWallet(db, log))
	e.GET("/query/wallet/events/:id", getWalletEvents(db, log))

	e.GET("/query/item/:id", getItem(db, log))
	e.GET("/query/item/events/:id", getItemEvents(db, log))
}
