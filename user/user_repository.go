package user

import (
	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	CreateUser(user *User) error
	GetUser(int) (User, error)
	UpdateUser(user *User, updateuser *UpdateUserRequest) error
	FindUser(map[string]interface{}) (User, error)
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) CreateUser(user *User) error {
	err := r.db.Create(&user).Error
	return err
}

func (r *Repository) GetUser(id int) (User, error) {
	user := User{}
	err := r.db.First(&user, id).Error
	return user, err
}
func (r *Repository) UpdateUser(user *User, updateuser *UpdateUserRequest) error {
	err := r.db.Model(&user).Updates(&updateuser).Error
	return err
}

func (r *Repository) FindUser(conditions map[string]interface{}) (User, error) {
	user := User{}
	err := r.db.Where(conditions).First(&user).Error
	return user, err
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
