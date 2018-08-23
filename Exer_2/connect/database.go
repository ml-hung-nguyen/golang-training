package connect

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

func NewConnect(db *gorm.DB) *gorm.DB {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpass, dbname)
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	fmt.Println("Connected Database")
	return db
}
