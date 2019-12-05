package main

import (
	"github.com/cyverse-de/de-stats/api"
	"github.com/cyverse-de/echo-middleware/redoc"
	"github.com/cyverse-de/de-stats/cron"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/cyverse-de/de-stats/docs"
)

func main() {
	db := cron.InitDB()
	defer db.Close()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(redoc.Serve(redoc.Opts{Title: "DE Stats API Documentation"}))

	e.GET("/", api.RootHandler)

	e.GET("/apps", api.AppsHandler)
	e.GET("/users", api.UsersHandler)
	e.GET("/jobs/submitted", api.JobsSubmittedHandler)
	e.GET("/jobs/status", api.JobsStatusHandler)
	e.GET("/logins", api.LoginCountHandler)
	e.GET("/logins/distinct", api.DistinceLoginCountHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
