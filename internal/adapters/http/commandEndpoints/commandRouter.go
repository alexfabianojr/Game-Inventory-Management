package commandEndpoints

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func StartCommandRouter(e *echo.Echo, db *sql.DB, log *zap.SugaredLogger) {
	e.POST("/command/inventory/create/:externalReference", createInventoryHandler(db, log))
	e.POST("/command/item/create/:externalReference/:inventoryId", createItemHandler(db, log))
}
