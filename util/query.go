package util

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

// ErrorResponse describes an error response for any endpoint.
type ErrorResponse struct {
	Description string `json:"description"`
}

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
	const (
		dateFormat = "20060102"
	)

	currentTime := time.Now()
	oneWeekAgo := time.Now().AddDate(0, 0, -7)

	startDate, err := StringQueryParam(ctx, "startDate", oneWeekAgo.Format(dateFormat))
	if err != nil {
		return "", "", ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	endDate, err := StringQueryParam(ctx, "endDate", currentTime.Format(dateFormat))
	if err != nil{
		return "", "", ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	return startDate, endDate, nil
}