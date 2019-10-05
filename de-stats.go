package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/pdrum/swagger-automation/docs"
)

type RootResponse struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

func rootHandler(ctx echo.Context) error {
	resp := RootResponse{
		Foo: "oof",
		Bar: "rab",
	}
	return ctx.JSON(http.StatusOK, resp)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", rootHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
