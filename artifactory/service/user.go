package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo) {
	// GET
	e.GET("/users", getAllUsers)
	e.GET("/user/:id", getUser)
	e.GET("/user/:id/settings", getUserSettings)

	// POST
	e.POST("/user", createUser)
	e.POST("/user/:id", updateUser)
	e.POST("/user/:id/settings", updateUserSettings)

	// DELETE
	e.DELETE("/user/:id", deleteUser)
}

func getAllUsers(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func getUser(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func getUserSettings(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func createUser(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func updateUser(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func updateUserSettings(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func deleteUser(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}
