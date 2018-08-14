package user

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type UserInterface interface {
	Find(id int, db *gorm.DB) error
	Create(db *gorm.DB) error
	Update(updatedUser *User, db *gorm.DB) error
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username" validate:"required,unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	FullName  string `json:"fullname" validate:"required"`
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

func (u *User) Create(db *gorm.DB) error {
	if db.Create(&u).RowsAffected < 1 {
		return errors.New("Something happen")
	}
	return nil
}

func (u *User) Update(updatedUser *User, db *gorm.DB) error {
	if db.Model(&u).Updates(&updatedUser).RowsAffected < 0 {
		return errors.New("Update fail")
	}
	return nil
}
