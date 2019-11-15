package util

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
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