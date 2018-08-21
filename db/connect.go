package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "1234"
	dbname   = "test_db"
)

var DB *gorm.DB

func Connect() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err = gorm.Open("postgres", psqlInfo)
	DB.LogMode(true)
	if err != nil {
		panic(err)
	}
}
