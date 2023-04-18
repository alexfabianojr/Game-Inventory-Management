package commandEndpoints

import (
	"database/sql"
	"errors"
	"game-inventory-management/internal/application/walletBusiness"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func addBalanceToWalletHandler(db *sql.DB, log *zap.SugaredLogger) echo.HandlerFunc {
	return func(c echo.Context) error {
		walletId, err := uuid.Parse(c.Param("walletId"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		test, err := strconv.ParseBool(c.QueryParam("test"))

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		request := new(AddBalanceRequest)
		err = c.Bind(request)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		_, err = walletBusiness.AddBalanceToWallet(walletId, request.Value, test, db, log)

		if err != nil {
			log.Error(err)
			return errors.New(err.Error())
		}

		return c.NoContent(http.StatusCreated)
	}
}

type AddBalanceRequest struct {
	Value int64 `json:"value"`
}
