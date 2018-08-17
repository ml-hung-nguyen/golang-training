package file

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type RepositoryInterface interface {
	FindUser(conditions map[string]interface{}) (User, error)
	CreateUser(user *User) error
	UpdateUser(user *User, updatedUser *User) error
	// CreatePost(user *User, content string) (Post, error)
	// GetPosts(conditions map[string]interface{}) ([]Post, error)
	// FindPost(id int) (Post, error)
	// UpdatePost(post *Post, content string) (Post, error)
	// DeletePost(post *Post) error
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
func (r *Repository) UpdateUser(user *User, updateUser *User) error {
	fmt.Println(updateUser)
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
