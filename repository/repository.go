package repository

import (
	"golang-training/model"
	"net/http"

	"github.com/go-playground/form"
	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	FindUser(map[string]interface{}) (model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(user *model.User) error

	DetailPost(map[string]interface{}) (model.Post, error)
	CreatePost(post *model.Post) error
	UpdatePost(post *model.Post) error
}

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func ParseForm(r *http.Request, i interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	decoder := form.NewDecoder()
	err = decoder.Decode(&i, r.Form)
	return err
}

func (r *Repository) FindUser(conditions map[string]interface{}) (model.User, error) {
	user := model.User{}
	err := r.DB.Where(conditions).First(&user).Error
	return user, err
}

func (r *Repository) CreateUser(user *model.User) error {
	err := r.DB.Create(&user).Error
	return err
}

func (r *Repository) UpdateUser(user *model.User) error {
	if err := r.DB.Model(&user).Update(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteUser(user *model.User) error {
	if err := r.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DetailPost(conditions map[string]interface{}) (model.Post, error) {
	post := model.Post{}
	err := r.DB.Where(conditions).First(&post).Error
	return post, err
}

func (r *Repository) CreatePost(post *model.Post) error {
	err := r.DB.Create(&post).Error
	return err
}
func (r *Repository) UpdatePost(post *model.Post) error {
	if err := r.DB.Model(&post).Update(&post).Error; err != nil {
		return err
	}
	return nil
}
