package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	// gorm.Model
	Id        int
	Username  string `gorm:"username"`
	FullName  string `gorm:"full_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Password  string `gorm:"password"`
}

type UserInterface interface {
	Detail(id int, db *gorm.DB) error
	Create(db *gorm.DB) error
	Update(userUpdate User, db *gorm.DB) error
}

func (user *User) Detail(id int, db *gorm.DB) error {
	result := db.Select("id, full_name, username, created_at, updated_at").Find(&user, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *User) Create(db *gorm.DB) error {
	user.Id = 0
	// if db.NewRecord(user) {
	// 	db.Create(&user)
	// }
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) Update(userUpdate User, db *gorm.DB) error {
	db.Model(&user).Update(&userUpdate)
	return nil
}
