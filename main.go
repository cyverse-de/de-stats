package main

import (
	"flag"
	"fmt"
	"github.com/cyverse-de/de-stats/api"
	"github.com/cyverse-de/de-stats/cron"
	"github.com/cyverse-de/de-stats/logs"
	_ "github.com/cyverse-de/de-stats/docs"
	"github.com/cyverse-de/echo-middleware/redoc"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/cyverse-de/configurate"
	"github.com/spf13/viper"
)

func main() {
	var (
		cfgPath = flag.String("config", "/etc/iplant/de/jobservices.yml", "The path to the config file")
		port = flag.String("port", "8080", "The port to listen on")
		debug = flag.String("debug", "off", "Turn on debug mode")
		err error
		cfg *viper.Viper
	)

	flag.Parse()

	logs.Init(debug)

	if *cfgPath == "" {
		//print error
		logs.Logger.Fatal("--config must not be the empty string")
	}

	if cfg, err = configurate.Init(*cfgPath); err != nil {
		logs.Logger.Fatal(err.Error())
	}

	dburi := cfg.GetString("db.uri")
	db, err := cron.InitDB(dburi)
	if err != nil {
		logs.Logger.Fatal(err.Error())
	}
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

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", *port)))
}
