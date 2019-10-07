package main

import (
	"github.com/cyverse-de/de-stats/api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/cyverse-de/de-stats/docs"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", api.RootHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
