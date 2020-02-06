package util

import (
	"fmt"
	"github.com/cyverse-de/de-stats/logs"
	"github.com/labstack/echo"
	"strconv"
	"time"
)

// ErrorResponse describes an error response for any endpoint.
type ErrorResponse struct {
	Description string `json:"description"`
}

const (
	dateFormat = "20060102"
)

// IntQueryParam extracts the value of an integer query parameter and performs range checking.
func IntQueryParam(ctx echo.Context, name string, defaultValue, minValue, maxValue int) (int, error) {

	// Get the query parameter value.
	valueStr := ctx.QueryParam(name)
	if valueStr == "" {
		return defaultValue, nil
	}

	// Parse the query parameter.
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		logs.Logger.Error(err)
		return 0, err
	}

	// Validate the query parameter.
	if value < minValue || value > maxValue {
		return 0, fmt.Errorf("'%s' query parameter (%d) must be in the range %d to %d", name, value, minValue, maxValue)
	}

	return value, nil
}

func StringQueryParam(ctx echo.Context, name string, defaultValue string) (string, error) {

	// Get the query parameter value.
	valueStr := ctx.QueryParam(name)
	if valueStr == "" {
		return defaultValue, nil
	}

	return valueStr, nil
}

func VerifyDateParameters(ctx echo.Context) (string, string, error) {
	currentTime := time.Now()
	oneWeekAgo := time.Now().AddDate(0, 0, -7)

	startDate, err := StringQueryParam(ctx, "startDate", oneWeekAgo.Format(dateFormat))
	if err != nil {
		logs.Logger.Error(err)
		return "", "", err
	}

	endDate, err := StringQueryParam(ctx, "endDate", currentTime.Format(dateFormat))
	if err != nil{
		logs.Logger.Error(err)
		return "", "", err
	}

	return startDate, endDate, nil
}