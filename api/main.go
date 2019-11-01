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


// RootHandler handles get requests to the / endpoint.
func RootHandler(ctx echo.Context) error {
	resp := RootResponse{
		Foo: "oof",
		Bar: "rab",
	}
	return ctx.JSON(http.StatusOK, resp)
}


