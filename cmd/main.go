package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	p := os.Getenv("HTTP_PORT")

	e.Logger.Fatal(e.Start(":" + p))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
