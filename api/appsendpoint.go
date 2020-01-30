package api

import (
	"database/sql"
	"github.com/cyverse-de/de-stats/cron"
	"github.com/cyverse-de/de-stats/logs"
	"github.com/cyverse-de/de-stats/util"
	"github.com/labstack/echo"
	"net/http"
)

type AppsParams struct {
	//The beginning of the time period of the response
	//
	//in: query
	//required: false
	//default: one week ago
	StartDate string

	//The end of the time period of the response
	//
	//in: query
	//required: false
	//default: today
	EndDate string

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

func BuildAppsHandler(db *sql.DB) func(echo.Context) error {
	return func(ctx echo.Context) error {
		startDate, endDate, err := util.VerifyDateParameters(ctx)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
		}

		amount, err := util.IntQueryParam(ctx, "count", 10, 1, 1000)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
		}

		apps, err := cron.GetTopApps(db, amount, startDate, endDate)

		if err != nil{
			logs.Logger.Error(err)
			return err
		}

		resp := AppsResponse{
			Count: len(apps),
			Apps:  apps,
		}
		return ctx.JSON(http.StatusOK, resp)
	}
}
