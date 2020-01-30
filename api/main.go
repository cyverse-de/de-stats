package api

import (
	"github.com/labstack/echo"
	"net/http"
)

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
	// Build the response.
	resp := RootResponse{
		Foo: "oof",
		Bar: "rab",
	}
	return ctx.JSON(http.StatusOK, resp)
}

