package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/smmir-cent/Simple-Backend/database"
)

type Public struct {
	DB database.DataBase
}

type CommentOutput struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Body   string `json:"body"`
}

type Comments struct {
	Movie    string           `json:"movie"`
	Comments []*CommentOutput `json:"comments"`
}

type MovieOutput struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rating      int    `json:"rating"`
}

type MoviesList struct {
	Movies []*MovieOutput `json:"movies"`
}

type MovieVote struct {
	Movie_id int `json:"movie_id"`
	Vote     int `json:"vote"`
}

type MovieComment struct {
	Movie_id     int    `json:"movie_id"`
	Comment_body string `json:"comment_body"`
}

func (public Public) GetComment(c echo.Context) error {

	if value := c.QueryParam("movie"); value != "" {
		intVar, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("can not cast id to int")
		}
		comments := public.DB.GetComments(intVar)

		var commentOutputs []*CommentOutput
		for _, element := range comments {
			commentOutputs = append(commentOutputs, &CommentOutput{Id: strconv.Itoa(element.Id), Author: strconv.Itoa(element.UserId), Body: element.Comment})
		}
		output := Comments{Movie: value, Comments: commentOutputs}
		return c.JSONPretty(http.StatusOK, output, "  ")
	}
	return c.NoContent(http.StatusNoContent)
}

func (public Public) GetMovies(c echo.Context) error {
	movies := public.DB.GetMovies()
	var MovieOutputs []*MovieOutput
	for _, element := range movies {
		MovieOutputs = append(MovieOutputs, &MovieOutput{Id: element.Id, Name: element.Name, Description: element.Description, Rating: int(element.Rating)})
	}
	output := MoviesList{Movies: MovieOutputs}
	return c.JSONPretty(http.StatusOK, output, "  ")
}

func (public Public) GetMovie(c echo.Context) error {
	log.Println("GET MOVIE")
	if value := c.Param("id"); value != "" {
		// log.Println(value)
		intVar, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("can not cast id to int")
		}
		movie := public.DB.GetMovie(intVar)
		if movie.Id != 0 {
			movieOutput := MovieOutput{Id: movie.Id, Name: movie.Name, Description: movie.Description, Rating: int(movie.Rating)}
			return c.JSONPretty(http.StatusOK, movieOutput, "  ")

		} else {
			//todo
			return c.NoContent(http.StatusNoContent)

		}
	}
	return c.NoContent(http.StatusNoContent)

}

func (public Public) MovieVote(c echo.Context) error {
	var Vote MovieVote
	if err := c.Bind(&Vote); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else {
		err := public.DB.InsertVote(database.Vote{UserId: 1, MovieId: Vote.Movie_id, Rating: float64(Vote.Vote)})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else {
			return c.NoContent(http.StatusNoContent)
		}
	}
}

func (public Public) CommentSubmit(c echo.Context) error {
	var Comment MovieComment
	if err := c.Bind(&Comment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else {
		err := public.DB.InsertComment(database.Comment{UserId: 1, MovieId: Comment.Movie_id, Comment: Comment.Comment_body, CreatedAt: time.Now().String()})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		} else {
			return c.NoContent(http.StatusNoContent)
		}
	}

}
