package user

import (
	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	FindUser(conditions map[string]interface{}) (User, error)
	CreateUser(user *User) error
	UpdateUser(updatedUser *User) (User, error)
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) FindUser(conditions map[string]interface{}) (User, error) {
	user := User{}
	// fmt.Println(r.DB)
	err := r.DB.Where(conditions).First(&user).Error
	// fmt.Println(user)
	return user, err
}

func (r *Repository) CreateUser(user *User) error {
	err := r.DB.Create(&user).Error
	return err
}

func (r *Repository) UpdateUser(updatedUser *User) (User, error) {
	var user User
	err := r.DB.Model(&user).Updates(&updatedUser).Error
	return user, err
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
