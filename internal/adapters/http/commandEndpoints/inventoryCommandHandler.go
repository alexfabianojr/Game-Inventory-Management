package commandEndpoints

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/application/inventoryBusiness"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func createInventoryHandler(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		externalReference := c.Param("externalReference")

		if externalReference == "" {
			message := "External reference not informed"
			log.Error(message)
			return errors.New(message)
		}

		_, err := inventoryBusiness.Create(externalReference, db, log)
		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.NoContent(http.StatusCreated)
	}
}
