package user

import (
	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	FindUser(conditions map[string]interface{}) (User, error)
	CreateUser(user *User) error
	UpdateUser(user *User, updatedUser *UpdateUserRequest) error
}
type Repository struct {
	DB *gorm.DB
}

//Create new user
func (r *Repository) CreateUser(user *User) error {
	err := r.DB.Create(&user).Error
	return err
}

//Update user
func (r *Repository) UpdateUser(user *User, updateUser *UpdateUserRequest) error {
	err := r.DB.Model(&user).Updates(&updateUser).Error
	return err
}

//Find user
func (r *Repository) FindUser(conditions map[string]interface{}) (User, error) {
	user := User{}
	err := r.DB.Where(conditions).First(&user).Error
	return user, err

}
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
