package api

import (
	"github.com/cyverse-de/de-stats/cron"
	"github.com/labstack/echo"
	"net/http"
)

type AppsParams struct {
	//The number of days to include in the response
	//
	//in: query
	//required: false
	//minimum: 0
	//maximum: 365
	//default: 7
	Days int

	//The number of apps to include in the response
	//
	//in: query
	//required: false
	//minimum: 1
	//maximum: 1000
	//default: 10
	Count int
}

type AppsResponse struct {
	Count	int `json:"count"`
	Apps	[]cron.App	`json:"apps"`
}

func AppsHandler(ctx echo.Context) error{
	db := cron.InitDB()
	amount := 10
	days := 100

	apps, err := cron.GetTopApps(db, amount, days)

	if err != nil{
		panic(err)
	}

	resp := AppsResponse{
		Count: amount,
		Apps:  apps,
	}
	return ctx.JSON(http.StatusOK, resp)

}