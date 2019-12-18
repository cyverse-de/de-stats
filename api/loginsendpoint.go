package api

import (
	"database/sql"
	"github.com/cyverse-de/de-stats/cron"
	"github.com/cyverse-de/de-stats/util"
	"github.com/labstack/echo"
	"net/http"
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

func BuildDistinctLoginCountHandler(db *sql.DB) func(echo.Context) error {
	return func(ctx echo.Context) error {
		startDate, endDate, err := util.VerifyDateParameters(ctx)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
		}

		login, err := cron.GetDistinctLoginCount(db, startDate, endDate)

		if err != nil {
			return err
		}

		resp := LoginsResponse{
			Count: login.Count,
		}

		return ctx.JSON(http.StatusOK, resp)
	}
}

func BuildLoginCountHandler(db *sql.DB) func(echo.Context) error {
	return func(ctx echo.Context) error {
		startDate, endDate, err := util.VerifyDateParameters(ctx)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
		}

		login, err := cron.GetLoginCount(db, startDate, endDate)

		if err != nil {
			return err
		}

		resp := LoginsResponse{
			Count: login.Count,
		}

		return ctx.JSON(http.StatusOK, resp)
	}
}