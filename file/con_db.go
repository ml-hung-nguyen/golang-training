package user

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

func InitDB() *sql.DB {
	config := configDB()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport], config[dbuser],
		config[dbpass], config[dbname])
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success")
	return db
}

func configDB() map[string]string {
	config := make(map[string]string)
	config[dbhost] = "localhost"
	config[dbport] = "5432"
	config[dbuser] = "postgres"
	config[dbpass] = "postgres"
	config[dbname] = "baitap-go1"
	return config
}
