package api

import (
	"github.com/cyverse-de/de-stats/cron"
	"net/http"
	"github.com/labstack/echo"
)

// RootResponse describes the response of the root endpoint.
type RootResponse struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

type AppsResponse struct {
	Count	int `json:"count"`
	Apps	[]cron.App	`json:"apps"`
}
// RootHandler handles get requests to the / endpoint.
func RootHandler(ctx echo.Context) error {
	resp := RootResponse{
		Foo: "oof",
		Bar: "rab",
	}
	return ctx.JSON(http.StatusOK, resp)
}

func AppsHandler(ctx echo.Context) error{
	db := cron.InitDB()
	amount := 10
	days := 100
	apps := cron.GetTopApps(db, amount, days)
	resp := AppsResponse{
		Count: amount,
		Apps:  apps,
	}
	return ctx.JSON(http.StatusOK, resp)

}
