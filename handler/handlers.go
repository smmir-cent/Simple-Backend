package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Public struct {
}

type Movie struct {
	Id          int
	Name        string
	Description string
	Rating      float64
}
type Vote struct {
	Id      int
	UserId  int
	MovieId int
	Rating  float64
}
type Comment struct {
	Id        int
	UserId    int
	MovieId   int
	CreatedAt string
	Comment   string
}
type User struct {
	Id       int
	Role     int
	UserName string
	Password string
}

func (public Public) GetComment(c echo.Context) error {
	if value := c.QueryParam("movie"); value != "" {
		log.Println(value)
		// TODO: DB select comments from movie where id == value
		// TODO: return in json with 200 status
		// {
		//   "movie": "string",
		//   "comments": [
		//     {
		//       "id": "Unknown Type: id",
		//       "author": "string",
		//       "body": "string"
		//     }
		//   ]
		// }

	}
	return c.NoContent(http.StatusNoContent)
}

func (public Public) GetMovies(c echo.Context) error {
	log.Println("GET MOVIES")
	// TODO: DB select from movie
	// TODO: return in json with 200 status
	// {
	//   "movies": [
	//     {
	//       "id": 0,
	//       "name": "string",
	//       "description": "string",
	//       "rating": 0
	//     }
	//   ]
	// }

	return c.NoContent(http.StatusNoContent)
}

func (public Public) GetMovie(c echo.Context) error {
	log.Println("GET MOVIE")
	if value := c.Param("id"); value != "" {
		log.Println(value)
		// TODO: DB select from movie where id = movies
		// TODO: return in json with 200 status
		// {
		//   "id": 0,
		//   "name": "string",
		//   "description": "string",
		//   "rating": 0
		// }

	}

	return c.NoContent(http.StatusNoContent)
}
