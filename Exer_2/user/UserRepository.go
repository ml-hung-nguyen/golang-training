package user

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepositoryInterface interface {
	DetailUser(user *User, con interface{}) error
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(user *User) error
}

// DetailUser get info user
func (repo *UserRepository) DetailUser(user *User, con interface{}) error {
	if con != nil {
		if err := repo.DB.Where(con).First(&user).Error; err != nil {
			return err
		}
	} else {
		if err := repo.DB.First(&user).Error; err != nil {
			return err
		}
	}

	return nil
}

// CreateUser create new user
func (repo *UserRepository) CreateUser(user *User) error {
	if repo.DB.NewRecord(user) {
		if err := repo.DB.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}

// UpdateUser update user
func (repo *UserRepository) UpdateUser(user *User) error {
	fmt.Println("OK", user)
	if err := repo.DB.Model(&user).Update(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) DeleteUser(user *User) error {
	if err := repo.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
