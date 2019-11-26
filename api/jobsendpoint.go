package api

import (
	"fmt"
	"github.com/cyverse-de/de-stats/cron"
	"github.com/cyverse-de/de-stats/util"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type JobsParams struct {
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

type JobsResponse struct {
	Count int	 	`json:"count"`
	Jobs []cron.Job `json:"jobs"`
}

func JobsSubmittedHandler(ctx echo.Context) error {
	const (
		dateFormat = "2006-01-02"
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

	db := cron.InitDB()
	jobs, err = cron.getSubmittedJobCounts(db, startDate, endDate)
	
	if err != nil {
		return err
	}

	count := len(jobs)
	resp := JobsResponse{
		Count: count,
		Jobs:  jobs,
	}

	return ctx.JSON(http.StatusOK, resp)

}