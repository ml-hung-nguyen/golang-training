package repository

import (
	"github.com/at-hungnguyen2/golang-training/model"
	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	CreateUser(string, string) (model.User, error)
	FindUser(map[string]interface{}) (model.User, error)
	// func CreatePost()
	// func GetPost()
	// func GetPosts()
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) CreateUser(username, fullname string) (model.User, error) {
	user := model.User{
		Username: username,
		FullName: fullname,
	}
	err := r.db.Create(&user).Error
	return user, err
}

func (r *Repository) FindUser(conditions map[string]interface{}) (model.User, error) {
	user := model.User{}
	err := r.db.Where(conditions).First(&user).Error
	return user, err
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
