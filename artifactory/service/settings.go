package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitSettingRoutes(e *echo.Echo) {

	//GET
	e.GET("/settings", getSettings)

	// POST
	e.POST("/settings", updateSettings)

}

func getSettings(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func updateSettings(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}
