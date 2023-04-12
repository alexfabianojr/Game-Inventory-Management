package httpCommand

import (
	"database/sql"
	"errors"
	logger "game-inventory-management/internal/adapters/log"
	"net/http"

	"github.com/labstack/echo"
)

func createInventoryHandler(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		log := logger.Get()
		defer logger.Get()
		externalReference := c.Param("externalReference")

		if externalReference == "" {
			message := "External reference not informed"
			log.Error(message)
			return errors.New(message)
		}

		return c.NoContent(http.StatusCreated)
	}
}
