package api

import (
	"github.com/cyverse-de/de-stats/cron"
	"github.com/cyverse-de/de-stats/util"
	"github.com/labstack/echo"
	"net/http"
)

type UsersParams struct {
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

	//The number of users to include in the response
	//
	//in: query
	//required: false
	//minimum: 1
	//maximum: 1000
	//default: 10
	Count int
}

type UsersResponse struct {
	Count	int `json:"count"`
	Users	[]cron.User	`json:"users"`
}

func UsersHandler(ctx echo.Context) error{
	startDate, endDate, err := util.VerifyDateParameters(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	amount, err := util.IntQueryParam(ctx, "count", 10, 1, 1000)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	db := cron.InitDB()
	users, err := cron.GetTopUsers(db, amount, startDate, endDate)

	if err != nil{
		return err
	}

	resp := UsersResponse{
		Count: len(users),
		Users:  users,
	}
	return ctx.JSON(http.StatusOK, resp)

}