package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/smmir-cent/Simple-Backend/handler"
)

func main() {
	e := echo.New()

	public := handler.Public{}
	// public
	e.GET("/comments", public.GetComment)
	e.GET("/movies", public.GetMovies)
	e.GET("/movie/:id", public.GetMovie)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}
