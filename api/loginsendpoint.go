package api

import (
	"github.com/cyverse-de/de-stats/cron"
	"github.com/cyverse-de/de-stats/util"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type LoginsParams struct {
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
}

type LoginsResponse struct {
	Count int 	`json:"count"`
}

func verifyLoginParameters(ctx echo.Context) (string, string, error) {
	const (
		dateFormat = "20060102"
	)

	currentTime := time.Now()
	oneWeekAgo := time.Now().AddDate(0, 0, -7)

	startDate, err := util.StringQueryParam(ctx, "startDate", oneWeekAgo.Format(dateFormat))
	if err != nil {
		return "", "", ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	endDate, err := util.StringQueryParam(ctx, "endDate", currentTime.Format(dateFormat))
	if err != nil{
		return "", "", ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	return startDate, endDate, nil
}

func DistinceLoginCountHandler(ctx echo.Context) error {

	startDate, endDate, err := verifyLoginParameters(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	db := cron.InitDB()
	login, err := cron.GetDistinctLoginCount(db, startDate, endDate)

	if err != nil {
		return err
	}

	resp := LoginsResponse{
		Count: login.Count,
	}

	return ctx.JSON(http.StatusOK, resp)

}

func LoginCountHandler(ctx echo.Context) error {
	startDate, endDate, err := verifyLoginParameters(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	db := cron.InitDB()
	login, err := cron.GetLoginCount(db, startDate, endDate)

	if err != nil {
		return err
	}

	resp := LoginsResponse{
		Count: login.Count,
	}

	return ctx.JSON(http.StatusOK, resp)
}