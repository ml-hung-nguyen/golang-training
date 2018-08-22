package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

type ConReponsitory struct {
	DB *gorm.DB
}
type ConReponseInterface interface {
	InitDB() *gorm.DB
}

func InitDB() *gorm.DB {
	config := configDB()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport], config[dbuser],
		config[dbpass], config[dbname])
	db, err := gorm.Open("postgres", psqlInfo)

	//db.LogMode(true)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
	return db
}

func configDB() map[string]string {
	config := make(map[string]string)
	config[dbhost] = "db"
	config[dbport] = "5432"
	config[dbuser] = "postgres"
	config[dbpass] = "postgres"
	config[dbname] = "reponse_go3"
	return config
}
