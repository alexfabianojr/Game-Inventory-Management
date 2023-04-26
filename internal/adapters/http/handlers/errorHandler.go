package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Message string `json:"message"`
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := err.Error()

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message.(string)
	}

	errorResponse := Error{Message: message}
	c.JSON(code, errorResponse)
}
