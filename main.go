package main

import (
	"github.com/cyverse-de/de-stats/api"
	"github.com/cyverse-de/de-stats/cron"
	_ "github.com/cyverse-de/de-stats/docs"
	"github.com/cyverse-de/echo-middleware/redoc"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	db := cron.InitDB()
	defer db.Close()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(redoc.Serve(redoc.Opts{Title: "DE Stats API Documentation"}))

	e.GET("/", api.RootHandler)

	e.GET("/apps", api.BuildAppsHandler(db))
	e.GET("/users", api.BuildUsersHandler(db))
	e.GET("/jobs/counts", api.BuildJobsHandler(db))
	e.GET("/logins", api.BuildLoginCountHandler(db))

	e.Logger.Fatal(e.Start(":8080"))
}
