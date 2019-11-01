package api

import (
	"fmt"
	"net/http"

	"github.com/cyverse-de/de-stats/util"
	"github.com/labstack/echo"
)

// swagger:parameters misc getRoot
type RootParams struct {
	// The number of days to include in the response
	//
	// in: query
	// required: false
	// minimum: 0
	// maximum: 365
	Days int
}

// RootResponse describes the response of the root endpoint.
type RootResponse struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

// ErrorResponse describes an error response for any endpoint.
type ErrorResponse struct {
	Description string `json:"description"`
}

// RootHandler handles get requests to the / endpoint.
func RootHandler(ctx echo.Context) error {

	// Extract the days query parameter.
	days, err := util.IntQueryParam(ctx, "days", 1, 0, 365)
	fmt.Println(days)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse{Description: err.Error()})
	}

	// Build the response.
	resp := RootResponse{
		Foo: "oof",
		Bar: "rab",
	}
	return ctx.JSON(http.StatusOK, resp)
}
