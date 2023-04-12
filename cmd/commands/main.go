package main

import (
	"game-inventory-management/internal/adapters/database/connection"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	logger := getLog()
	defer getLog()
	logger.Info("Game Inventory Management - Commands")

	dbConnection := connection.Connection{
		Host:     "localhost",
		Port:     "5433",
		DbName:   "game_inventory_management",
		DbUser:   "postgres",
		Password: "123456",
	}

	connection.DatabaseConnection(dbConnection)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	healthEndpoint(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func getLog() *zap.SugaredLogger {
	log, _ := zap.NewProduction()
	return log.Sugar()
}

func healthEndpoint(e *echo.Echo) {
	response := func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	}
	e.GET("/", response)
}
