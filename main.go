package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})
	if err := e.Start("127.0.0.1:8080"); err != nil {
		log.Fatalf("can not start echo http srver %s", err)
	}
}
