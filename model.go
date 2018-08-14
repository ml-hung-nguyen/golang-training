package main

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type UserInterface interface {
	Detail(id int, db *gorm.DB) error
	Create(db *gorm.DB) error
	Update(userUpdate User, db *gorm.DB) error
}

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Fullname  string `json:"fullname"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (user *User) Create(db *gorm.DB) error {
	if db.NewRecord(user) {
		db.Create(&user)
	}
	return nil
}

func (user *User) Detail(id int, db *gorm.DB) error {
	if db.First(&user, id).RecordNotFound() {
		return errors.New("User not Found")
	}

	return nil
}

func (user *User) Update(userUpdate User, db *gorm.DB) error {
	db.Model(&user).Update(&userUpdate)
	return nil
}
