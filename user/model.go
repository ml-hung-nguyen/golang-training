package user

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type UserInterface interface {
	Find(id int, db *gorm.DB) error
	Create(user *User, db *gorm.DB) error
	Update(user *User, updatedUser *User, db *gorm.DB) error
}

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	FullName  string
}

type JwtToken struct {
	Token string `json:"token"`
}

type UserRequest struct {
	ID       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username" validate:"required,unique"`
	Password string `json:"password" form:"password"`
	FullName string `json:"fullname" form:"fullname"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	FullName string `json:"fullname"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (u *User) Find(id int, db *gorm.DB) error {
	u.ID = 0
	if db.Where("id = ?", id).First(&u).RecordNotFound() {
		return errors.New("Record not found")
	}
	return nil
}

func (u *User) Create(user *User, db *gorm.DB) error {
	if db.Create(&user).RowsAffected < 1 {
		return errors.New("Something happen")
	}
	return nil
}

func (u *User) Update(user *User, updatedUser *User, db *gorm.DB) error {
	if db.Model(&user).Updates(&updatedUser).RowsAffected < 0 {
		return errors.New("Update fail")
	}
	return nil
}
