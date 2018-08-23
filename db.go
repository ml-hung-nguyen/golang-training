package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func NewConnect() *gorm.DB {
	var err error
	DB, err = gorm.Open("postgres",
		"user=postgres port=5432 dbname=exercise password=root sslmode=disable")

	DB.LogMode(true)
	if err != nil {
		panic(err)
	}
	return DB
}
