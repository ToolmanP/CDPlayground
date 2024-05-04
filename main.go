package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Foo struct {
	Msg string `json:"msg"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &Foo{
			Msg: "Hello World",
		})
	})
	e.Logger.Fatal(e.Start(":10086"))
}
