package user

import (
	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	FindUser(conditions map[string]interface{}) (User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) FindUser(conditions map[string]interface{}) (User, error) {
	user := User{}
	err := r.DB.Where(conditions).First(&user).Error
	return user, err
}

func (r *Repository) CreateUser(user *User) error {
	err := r.DB.Create(&user).Error
	return err
}

func (r *Repository) UpdateUser(user *User) error {
	// fmt.Println(user)
	err := r.DB.Model(&user).Update(&user).Error
	return err
}
