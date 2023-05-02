package queryEndpoint

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/application/inventoryBusiness"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func getInventory(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		externalReference, err := uuid.Parse(c.Param("externalReference"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		inventory, err := inventoryBusiness.GetInventory(externalReference, db, log)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.JSON(http.StatusOK, inventory)
	}
}
