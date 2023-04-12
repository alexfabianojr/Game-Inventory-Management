package httpCommand

import (
	"database/sql"

	"github.com/labstack/echo"
)

func StartCommandRouter(e *echo.Echo, db *sql.DB) {
	e.POST("/command/inventory/create/:externalReference", createInventoryHandler(db))
}
