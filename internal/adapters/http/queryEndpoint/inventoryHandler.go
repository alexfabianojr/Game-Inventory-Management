package queryEndpoint

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/application/inventoryBusiness"
	projection "game-inventory-management/internal/application/projectionBusiness"
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

func GetInventoryWithWallet(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		inventory, err := projection.GetInventoryWithWalletByInventoryId(id, db, log)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.JSON(http.StatusOK, inventory)
	}
}

func GetInventoryWithItems(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		inventoryId, err := uuid.Parse(c.Param("inventory_id"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		inventory, err := projection.GetInventoryWithItems(inventoryId, db, log)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.JSON(http.StatusOK, inventory)
	}
}

func GetInventoryWithWalletAndItems(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		inventoryId, err := uuid.Parse(c.Param("inventory_id"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		inventory, err := projection.GetInventoryWithWalletAndItems(inventoryId, db, log)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.JSON(http.StatusOK, inventory)
	}
}
