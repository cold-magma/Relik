package service

import (
	"crypto/subtle"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitArtifactoryRoutes(e *echo.Echo) {

	//GET
	e.GET("/artifact", getAllArtifacts)
	e.GET("/artifact/search/{query}", searchArtifacts)
	e.GET("/artifact/count", getArtifactCount)
	e.GET("/artifact/:id", getArtifact)
	e.GET("/artifact/recent", getRecentArtifacts)

	//POST
	e.POST("/artifact", createArtifact)
	e.POST("/artifact/:id", updateArtifact)

	// DELETE
	e.DELETE("/artifact/:id", deleteArtifact)

	// artifact get and push
	g := e.Group("/artifact/java", middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			return true, nil
		}
		return false, nil
	}))
	g.POST("", createArtifact)
	g.GET("/:group/:artifact/:version", getArtifactDependency)
}

// GET
func getAllArtifacts(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func searchArtifacts(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func getArtifactCount(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func getArtifact(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, c.Param("id"), " ")
}

func getRecentArtifacts(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func getArtifactDependency(c echo.Context) error {
	fmt.Println(c.Param("group"))
	fmt.Println(c.Param("artifact"))
	fmt.Println(c.Param("version"))
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

// POST
func updateArtifact(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func createArtifact(c echo.Context) error {
	// get file
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

//DELETE

func deleteArtifact(c echo.Context) error {
	fmt.Println("hello world")
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}
