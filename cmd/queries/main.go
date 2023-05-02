package main

import (
	"game-inventory-management/internal/adapters/database/connection"
	"game-inventory-management/internal/adapters/http/handlers"
	"game-inventory-management/internal/adapters/http/queryEndpoint"
	logger "game-inventory-management/internal/adapters/log"
	"game-inventory-management/internal/adapters/properties"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log := logger.Get()
	defer logger.Get()
	log.Info("Game Inventory Management - Query")

	props, err := properties.Get(log)

	if err != nil {
		panic("Cannot load system properties")
	}

	db, err := connection.DatabaseConnection(props.Database, log)
	if err != nil {
		log.Fatal("Database error", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	startAliveEndpoint(e)
	queryEndpoint.StartQueryRouter(e, db, log)

	e.Logger.Fatal(e.Start(":1323"))
}

func startAliveEndpoint(e *echo.Echo) {
	response := func(c echo.Context) error {
		return c.String(http.StatusOK, "Game Inventory Management - Query: OK")
	}
	e.GET("/", response)
}
