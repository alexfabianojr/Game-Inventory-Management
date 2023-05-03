package queryEndpoint

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"go.uber.org/zap"
)

func StartQueryRouter(e *echo.Echo, db *sql.DB, log *zap.SugaredLogger) {
	e.GET("/query/inventory/:externalReference", getInventory(db, log))
	e.GET("/query/inventory/summary/total/:id", getInventory(db, log))
	e.GET("/query/inventory/summary/wallet/:id", getInventory(db, log))
	e.GET("/query/inventory/summary/wallet/items/:id", getInventory(db, log))
	e.GET("/query/inventory/summary/items/:id", getInventory(db, log))

	e.GET("/query/wallet/:id", getInventory(db, log))
	e.GET("/query/wallet/events/:id", getInventory(db, log))

	e.GET("/query/item/:id", getInventory(db, log))
	e.GET("/query/item/events/:id", getInventory(db, log))
}
