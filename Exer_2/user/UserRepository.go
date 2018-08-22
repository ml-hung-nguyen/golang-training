package user

import (
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
		err := repo.DB.Where(con).First(&user).Error
		return err
	}
	err := repo.DB.First(&user).Error
	return err
}

// CreateUser create new user
func (repo *UserRepository) CreateUser(user *User) error {
	var err error
	if repo.DB.NewRecord(user) {
		err = repo.DB.Create(&user).Error
	}
	return err
}

// UpdateUser update user
func (repo *UserRepository) UpdateUser(user *User) error {
	err := repo.DB.Model(&user).Update(&user).Error
	return err
}

func (repo *UserRepository) DeleteUser(user *User) error {
	err := repo.DB.Delete(&user).Error
	return err
}
