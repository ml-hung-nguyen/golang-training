package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "postgres"
	dbpass = "root"
	dbname = "postgres"
)

func newConnect() {
	var err error
	psql := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpass, dbname)
	db, err = gorm.Open("postgres", psql)
	if err != nil {
		panic(err)
	}
	// db.LogMode(true)
}
