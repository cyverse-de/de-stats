package api

import (
	"fmt"
	"github.com/cyverse-de/de-stats/cron"
	"github.com/labstack/echo"
	"net/http"
	"github.com/cyverse-de/de-stats/util"
	"time"
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

func AppsHandler(ctx echo.Context) error{
	const (
		dateFormat = "20060102"
	)

	currentTime := time.Now()
	oneWeekAgo := time.Now().AddDate(0, 0, -7)

	startDate, err := util.StringQueryParam(ctx, "startDate", oneWeekAgo.Format(dateFormat))
	fmt.Println(startDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	endDate, err := util.StringQueryParam(ctx, "endDate", currentTime.Format(dateFormat))
	fmt.Println(endDate)
	if err != nil{
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}
	eDate, err := time.Parse(dateFormat, endDate)
	eDate.AddDate(0, 0, 1)
	endDate = eDate.Format(dateFormat)


	amount, err := util.IntQueryParam(ctx, "count", 10, 1, 1000)
	fmt.Println(amount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	db := cron.InitDB()
	apps, err := cron.GetTopApps(db, amount, startDate, endDate)

	if err != nil{
		return err
	}

	resp := AppsResponse{
		Count: len(apps),
		Apps:  apps,
	}
	return ctx.JSON(http.StatusOK, resp)

}