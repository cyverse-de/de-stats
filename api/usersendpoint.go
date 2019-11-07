package api

import (
	"fmt"
	"github.com/cyverse-de/de-stats/cron"
	"github.com/labstack/echo"
	"net/http"
	"github.com/cyverse-de/de-stats/util"
)

type UsersParams struct {
	//The number of days to include in the response
	//
	//in: query
	//required: false
	//minimum: 0
	//maximum: 365
	//default: 7
	Days int

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

	days, err := util.IntQueryParam(ctx, "days", 1, 0, 365)
	fmt.Println(days)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}
	db := cron.InitDB()

	amount, err := util.IntQueryParam(ctx, "count", 10, 1, 1000)
	fmt.Println(amount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	users, err := cron.GetTopUsers(db, amount, days)

	if err != nil{
		return err
	}

	resp := UsersResponse{
		Count: amount,
		Users:  users,
	}
	return ctx.JSON(http.StatusOK, resp)

}