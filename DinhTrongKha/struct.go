package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int    `gorm:"id" form:"-"`
	Username  string `gorm:"username" form:"username"`
	FullName  string `gorm:"full_name" form:"fullname"`
	Password  string `gorm:"password" form:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserInterface interface {
	Detail(userR *User, db *gorm.DB) error
	Create(userR *User, db *gorm.DB) error
	Update(userR *User, db *gorm.DB) error
}

// Create create new user
func (user *User) Create(userR *User, db *gorm.DB) error {
	if db.NewRecord(userR) {
		if err := db.Create(&userR).Error; err != nil {
			return err
		}
	}
	return nil
}

// Detail get info user
func (user *User) Detail(userR *User, db *gorm.DB) error {
	result := db.First(&userR)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Update update user
func (user *User) Update(userR *User, db *gorm.DB) error {
	if err := db.Model(&userR).Update(&userR).Error; err != nil {
		return err
	}

	return nil
}

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	password string `json:"-"`
}
