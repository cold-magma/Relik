package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitAuthenticationRoutes(e *echo.Echo) {
	// POST
	e.POST("/login", login)
}

func login(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}
