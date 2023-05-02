package queryEndpoint

import (
	"database/sql"

	"github.com/labstack/echo/v4"

	"go.uber.org/zap"
)

func StartQueryRouter(e *echo.Echo, db *sql.DB, log *zap.SugaredLogger) {
	e.GET("/command/inventory/:externalReference", getInventory(db, log))
}
