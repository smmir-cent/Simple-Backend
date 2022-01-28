package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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
type DataBase struct {
	Connection *sql.DB
}

func (db *DataBase) Connect() {
	var err error
	db.Connection, err = sql.Open("mysql", "root:Password123#@!@tcp(127.0.0.1:3306)/hw4")
	if err != nil {
		panic(err)
	}
	// defer db.Connection.Close()
}

func (db *DataBase) GetComments(id int) []Comment {

	res, err := db.Connection.Query("SELECT * FROM comment WHERE movieId = ?", id)
	if err != nil {
		return nil
	}
	var comments []Comment
	for res.Next() {
		var comment Comment
		err := res.Scan(&comment.Id, &comment.UserId, &comment.MovieId, &comment.Comment, &comment.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}
	return comments
}

func (db *DataBase) GetMovies() []Movie {

	res, err := db.Connection.Query("SELECT * FROM movie")
	var movies []Movie
	if err != nil {
		return nil
	}
	for res.Next() {
		var movie Movie
		err := res.Scan(&movie.Id, &movie.Name, &movie.Description, &movie.Rating)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
		log.Printf("%v\n", movie)
	}
	return movies
}

func (db *DataBase) GetMovie(id int) Movie {
	var movie Movie
	err2 := db.Connection.QueryRow("SELECT * FROM movie WHERE id = ?", id).Scan(&movie.Id, &movie.Name, &movie.Description, &movie.Rating)
	if err2 != nil {
		log.Println(err2)
	}

	return movie
}

func (db *DataBase) InsertVote(v Vote) error {
	sqlStatement := "INSERT INTO vote (userId, movieId, rating) VALUES (?, ?, ?)"
	_, err := db.Connection.Exec(sqlStatement, v.UserId, v.MovieId, v.Rating)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (db *DataBase) InsertComment(c Comment) error {
	sqlStatement := "INSERT INTO comment (userId,  movieId, comment , createdAt) VALUES (?, ?, ? , ?)"
	_, err := db.Connection.Exec(sqlStatement, c.UserId, c.MovieId, c.Comment, c.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (db *DataBase) InsertMovie(movie Movie) error {

	sqlStatement := `INSERT INTO movie (name, description) VALUES (?, ?)`
	_, err := db.Connection.Exec(sqlStatement, movie.Name, movie.Description)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (db *DataBase) EditMovie(id int, name string, description string) error {
	sqlStatement := "UPDATE movie SET name = ?, description= ? WHERE id = ?"
	_, err := db.Connection.Exec(sqlStatement, name, description, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (db *DataBase) EditComment(id int, approved bool) error {
	sqlStatement := "UPDATE comment SET approved = ? WHERE id = ?"
	_, err := db.Connection.Exec(sqlStatement, approved, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (db *DataBase) DeleteMovie(id int) error {
	sqlStatement := "DELETE FROM movie WHERE id=?"
	_, err := db.Connection.Exec(sqlStatement, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (db *DataBase) DeleteComment(id int) error {
	sqlStatement := "DELETE FROM comment WHERE id=?"
	_, err := db.Connection.Exec(sqlStatement, id)
	if err != nil {
		log.Println(err)
	}
	return err
}
