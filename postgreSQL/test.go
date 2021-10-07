package main

import (
	"database/sql"

	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Lakshmi@123"
	dbname   = "sireesha"
)

var db *sql.DB
var err error

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	stmt, err := db.Prepare("INSERT INTO posts(id,title) VALUES($1,$2)")
	if err != nil {

		fmt.Println(err.Error())
	}
	_, err = stmt.Exec("1", "2")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)

	}

}
