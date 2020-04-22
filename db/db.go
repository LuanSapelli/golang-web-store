package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Dbconnector() *sql.DB {
	connection := "user=postgres dbname=alura_loja password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
