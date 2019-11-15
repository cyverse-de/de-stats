package api

import (
	"fmt"
	"github.com/cyverse-de/de-stats/cron"
	"github.com/labstack/echo"
	"net/http"
	"github.com/cyverse-de/de-stats/util"
	"time"
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

	currentTime := time.Now()
	oneWeekAgo := time.Now().AddDate(0, 0, -7)

	startDate, err := util.StringQueryParam(ctx, "startDate", oneWeekAgo.Format("00/00/0000"))
	fmt.Println(startDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	endDate, err := util.StringQueryParam(ctx, "endDate", currentTime.Format("00/00/0000"))
	fmt.Println(endDate)
	if err != nil{
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	amount, err := util.IntQueryParam(ctx, "count", 10, 1, 1000)
	fmt.Println(amount)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	db := cron.InitDB()
	users, err := cron.GetTopUsers(db, amount, startDate, endDate)

	if err != nil{
		return err
	}

	resp := UsersResponse{
		Count: amount,
		Users:  users,
	}
	return ctx.JSON(http.StatusOK, resp)

}