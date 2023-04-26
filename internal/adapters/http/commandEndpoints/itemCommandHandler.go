package commandEndpoints

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/application/itemBusiness"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func createItemHandler(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		externalReference, err := uuid.Parse(c.Param("externalReference"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		inventoryId, err := uuid.Parse(c.Param("inventoryId"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		test, err := strconv.ParseBool(c.QueryParam("test"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		_, err = itemBusiness.Create(inventoryId, externalReference, test, db, log)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.NoContent(http.StatusCreated)
	}
}

func itemExchangesBetweenInventories(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		requestBody := new(itemBusiness.ItemExchanges)
		err := c.Bind(requestBody)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		err = itemBusiness.ItemExchangesBetweenInventories(*requestBody, db, log)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.NoContent(http.StatusCreated)
	}
}
