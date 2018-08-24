package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Init() *gorm.DB {
	var err error
	DB, err = gorm.Open("postgres", "port=5432 host=db user=postgres dbname=crud_example password=sa sslmode=disable ")
	if err != nil {
		panic(err)
	}
	return DB
}
