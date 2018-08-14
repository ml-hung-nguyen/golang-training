package main

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	FullName  string `json:"fullname"`
}

type UserInterface interface {
	Detail(id int, db *gorm.DB) error
	Create(db *gorm.DB) error
	Update(userUpdate User, db *gorm.DB) error
}

// Create create new user
func (user *User) Create(db *gorm.DB) error {
	if db.NewRecord(user) {
		db.Create(&user)
	}
	return nil
}

// Detail get info user
func (user *User) Detail(id int, db *gorm.DB) error {
	if db.First(&user, id).RecordNotFound() {
		return errors.New("User not Found")
	}

	return nil
}

// Update update user
func (user *User) Update(userUpdate User, db *gorm.DB) error {
	// user.Id = id
	db.Model(&user).Update(&userUpdate)
	return nil
}
