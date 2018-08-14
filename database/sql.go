package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() *sql.DB {
	var err error
	DB, err = sql.Open("postgres", "user=postgres dbname=crud_example password=sa sslmode=disable ")
	if err != nil {
		panic(err)
	}
	return DB
}
