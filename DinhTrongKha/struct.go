package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int        `gorm:"id"`
	Username  string     `gorm:"username"`
	FullName  string     `gorm:"full_name"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}

type UserInterface interface {
	Detail(id int, db *gorm.DB) error
	Create(db *gorm.DB) error
	Update(userUpdate User, db *gorm.DB) error
}

// Create create new user
func (user *User) Create(db *gorm.DB) error {
	user.Id = 0
	if db.NewRecord(user) {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}

// Detail get info user
func (user *User) Detail(id int, db *gorm.DB) error {
	user.Id = 0
	result := db.First(&user, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Update update user
func (user *User) Update(userUpdate User, db *gorm.DB) error {
	db.Model(&user).Update(&userUpdate)
	return nil
}
