package queryEndpoint

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/application/walletBusiness"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func getWallet(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		inventory, err := walletBusiness.GetWallet(id, db, log)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.JSON(http.StatusOK, inventory)
	}
}

func getWalletEvents(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		inventory, err := walletBusiness.GetWalletEvents(id, db, log)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.JSON(http.StatusOK, inventory)
	}
}
