package main

import (
	"game-inventory-management/internal/adapters/database/connection"
	"game-inventory-management/internal/adapters/properties"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func main() {
	log := getLog()
	defer getLog()
	log.Info("Game Inventory Management - Commands")

	props, err := properties.Get(log)

	if err != nil {
		panic("Cannot load system properties")
	}

	connection.DatabaseConnection(props.Database, log)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	startAliveEndpoint(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func getLog() *zap.SugaredLogger {
	log, _ := zap.NewProduction()
	return log.Sugar()
}

func startAliveEndpoint(e *echo.Echo) {
	response := func(c echo.Context) error {
		return c.String(http.StatusOK, "Game Inventory Management - Commands: OK")
	}
	e.GET("/", response)
}
