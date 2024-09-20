package main

import (
	artifactory "service"
	authentication "service"
	files "service"
	settings "service"
	user "service"

	"github.com/labstack/echo/v4"
)

// main function
func main() {
	e := echo.New()

	artifactory.InitArtifactoryRoutes(e)
	authentication.InitAuthenticationRoutes(e)
	user.InitUserRoutes(e)
	files.InitFileRoutes(e)
	settings.InitSettingRoutes(e)

	e.Logger.Fatal(e.Start(":8000"))
}
