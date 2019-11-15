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
	apps, err := cron.GetTopApps(db, amount, startDate, endDate)

	if err != nil{
		return err
	}

	resp := AppsResponse{
		Count: amount,
		Apps:  apps,
	}
	return ctx.JSON(http.StatusOK, resp)

}