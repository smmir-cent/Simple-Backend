package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/smmir-cent/Simple-Backend/database"
	"github.com/smmir-cent/Simple-Backend/handler"
)

func main() {
	e := echo.New()
	db := database.DataBase{}
	db.Connect()

	public := handler.Public{DB: db}
	// public
	e.GET("/comments", public.GetComment)
	e.GET("/movies", public.GetMovies)
	e.GET("/movie/:id", public.GetMovie)
	// movie := database.Movie{
	// 	Id:          1,
	// 	Name:        "succession",
	// 	Description: "perfect",
	// 	Rating:      4.2,
	// }
	// db.InsertMovie(movie)
	// user
	e.POST("/user/vote", public.MovieVote)
	e.POST("/user/comment", public.CommentSubmit)
	if err := e.Start("0.0.0.0:8080"); err != nil {
		log.Fatal(err)
	}
}

/*

CREATE TABLE movie ( id integer, name varchar(32) , description varchar(32) , rating double );
CREATE TABLE vote ( id integer NOT NULL AUTO_INCREMENT, userId integer , movieId integer , rating double,PRIMARY KEY (id) );
CREATE TABLE comment ( id integer NOT NULL AUTO_INCREMENT, userId integer , movieId integer, comment varchar(256) ,approved BOOLEAN DEFAULT false, createdAt varchar(256),PRIMARY KEY (id)  );
CREATE TABLE user ( id integer, role integer , username varchar(256) , password varchar(256) );
*/
