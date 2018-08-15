package user

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

func InitDB() *gorm.DB {
	config := configDB()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport], config[dbuser],
		config[dbpass], config[dbname])
	db, err := gorm.Open("postgres", psqlInfo)
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
