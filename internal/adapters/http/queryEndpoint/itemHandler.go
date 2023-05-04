package queryEndpoint

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/application/itemBusiness"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func getItem(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		inventory, err := itemBusiness.GetItem(id, db, log)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.JSON(http.StatusOK, inventory)
	}
}

func getItemEvents(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		inventory, err := itemBusiness.GetItemEvents(id, db, log)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.JSON(http.StatusOK, inventory)
	}
}
