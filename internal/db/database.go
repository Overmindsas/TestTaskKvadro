package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

func DbConnect() (*sql.DB, error) {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:8000)/kvadodb?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db.Ping())

	return &sql.DB{}, nil

}

func (sq *Repo) DbFindAuthor() {
	res, err := sq.db.Query("SELECT book FROM BookAuthor WHERE author = 'King'")
	if err != nil {
		log.Println(err)
	}
	defer res.Close()

	for res.Next() {
		var book BookRepo
		err := res.Scan(&book.Book)
		if err != nil {
			log.Println(err)
		}

		fmt.Println(book)
	}

}

func (sq *Repo) DbFindBook() {

}
