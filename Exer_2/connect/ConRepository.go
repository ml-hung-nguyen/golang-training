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

type ConnectRepository struct {
	DB *gorm.DB
}

type ConnectRepositoryInterface interface {
	ConnectRepo() error
}

func (repo *ConnectRepository) ConnectRepo() error {
	var err error
	psql := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpass, dbname)
	repo.DB, err = gorm.Open("postgres", psql)
	if err != nil {
		panic(err)
	}
	// repo.db.Close()
	// repo.db.LogMode(true)
	return nil
}

func NewConRepo(db *gorm.DB) *ConnectRepository {
	return &ConnectRepository{
		DB: db,
	}
}

func NewConnect(db *gorm.DB) *gorm.DB {
	var err error
	psql := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		dbhost, dbport, dbuser, dbpass, dbname)
	db, err = gorm.Open("postgres", psql)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	fmt.Println("Connected Database")
	return db
}
