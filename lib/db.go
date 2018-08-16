package lib

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func NewConnect() *sql.DB {
	var err error
	DB, err = sql.Open("postgres", "user=postgres dbname=exercise password=root sslmode=disable")
	if err != nil {
		panic(err)
	}
	return DB
}
