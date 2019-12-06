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
	Jobs []cron.JobStats `json:"jobs"`
}

func verifyJobsParameters(ctx echo.Context) (string, string, error) {
	const (
		dateFormat = "20060102"
	)

	currentTime := time.Now()
	oneWeekAgo := time.Now().AddDate(0, 0, -7)

	startDate, err := util.StringQueryParam(ctx, "startDate", oneWeekAgo.Format(dateFormat))
	fmt.Println(startDate)
	if err != nil {
		return "", "", ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	endDate, err := util.StringQueryParam(ctx, "endDate", currentTime.Format(dateFormat))
	fmt.Println(endDate)
	if err != nil{
		return "", "", ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	return startDate, endDate, nil
}

func JobsSubmittedHandler(ctx echo.Context) error {

	startDate, endDate, err := verifyJobsParameters(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	db := cron.InitDB()
	jobs, err := cron.GetSubmittedJobCounts(db, startDate, endDate)
	
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

func JobsStatusHandler(ctx echo.Context) error {
	startDate, endDate, err := verifyJobsParameters(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	db := cron.InitDB()
	jobs, err := cron.GetJobStatusCounts(db, startDate, endDate)

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
