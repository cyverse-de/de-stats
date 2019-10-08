package main

import (
	"github.com/cyverse-de/de-stats/api"
	"github.com/cyverse-de/de-stats/redoc"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/cyverse-de/de-stats/docs"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(redoc.Serve(redoc.Opts{Title: "DE Stats API Documentation"}))

	e.GET("/", api.RootHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
