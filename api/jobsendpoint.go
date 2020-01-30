package api

import (
	"database/sql"
	"github.com/cyverse-de/de-stats/cron"
	"github.com/cyverse-de/de-stats/logs"
	"github.com/cyverse-de/de-stats/util"
	"github.com/labstack/echo"
	"net/http"
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

func BuildJobsHandler(db *sql.DB) func(echo.Context) error {
	return func(ctx echo.Context) error {
		startDate, endDate, err := util.VerifyDateParameters(ctx)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
		}

		jobs, err := cron.GetJobCounts(db, startDate, endDate)

		if err != nil {
			logs.Logger.Error(err)
			return err
		}

		count := len(jobs)
		resp := JobsResponse{
			Count: count,
			Jobs:  jobs,
		}

		return ctx.JSON(http.StatusOK, resp)
	}
}
