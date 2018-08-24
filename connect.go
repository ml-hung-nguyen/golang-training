package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "postgres"
	dbpass = "root"
	dbname = "test_db"
)

func Connect() *gorm.DB {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpass, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	return db
}
