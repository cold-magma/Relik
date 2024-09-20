// for uploading downloading artifacts from the relik UI
package service

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func InitFileRoutes(e *echo.Echo) {
	// GET
	e.GET("/download", downloadFile)

	//POST
	e.POST("/upload", uploadFile)
}

func downloadFile(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, "hello world", " ")
}

func uploadFile(c echo.Context) error {
	// Get file
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("Did not revieve file")
		panic(err)
		return err
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println("Did not open file")
		panic(err)
		return err
	}
	defer src.Close()

	dst, err := os.Create("uploads/" + file.Filename)
	if err != nil {
		fmt.Println("Did not create file")
		panic(err)
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println("Did not copy file")
		panic(err)
		return err
	}
	fmt.Println("Uploaded file")
	return c.HTML(http.StatusOK, "FIle uploaded")
}
