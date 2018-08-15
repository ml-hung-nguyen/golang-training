package user

import (
	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	FindUser(conditions map[string]interface{}) (User, error)
	CreateUser(user *User) error
	UpdateUser(user *User, updatedUser *User) error
	CreatePost(user *User, content string) (Post, error)
	GetPosts(conditions map[string]interface{}) ([]Post, error)
	FindPost(id int) (Post, error)
	UpdatePost(post *Post, content string) (Post, error)
	DeletePost(post *Post) error
}

type Repository struct {
	DB *gorm.DB
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

func (r *Repository) UpdateUser(user *User, updatedUser *User) error {
	err := r.DB.Model(&user).Updates(&updatedUser).Error
	return err
}

func (r *Repository) CreatePost(user *User, content string) (Post, error) {
	post := Post{
		IdUser:  user.ID,
		Content: content,
	}
	err := r.DB.Create(&post).Error
	return post, err
}

func (r *Repository) GetPosts(conditions map[string]interface{}) ([]Post, error) {
	posts := []Post{}
	err := r.DB.Where(conditions).Find(&posts).Error
	return posts, err
}

func (r *Repository) FindPost(id int) (Post, error) {
	post := Post{}
	err := r.DB.Find(&post, id).Error
	return post, err
}

func (r *Repository) UpdatePost(post *Post, content string) (Post, error) {
	post.Content = content
	err := r.DB.Save(&post).Error
	return *post, err
}

func (r *Repository) DeletePost(post *Post) error {
	err := r.DB.Delete(&post).Error
	return err
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}
