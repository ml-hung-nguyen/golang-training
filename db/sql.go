package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func NewConnect() {
	// dbms := "postgres"
	// user := "postgres"
	// password := ""
	// dbname := "golang"
	// connecStr := "user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=disable"
	var err error
	DB, err = gorm.Open("postgres",
		"user=postgres dbname=golang password=gwp sslmode=disable")

	DB.LogMode(true)
	if err != nil {
		panic(err)
	}
}
