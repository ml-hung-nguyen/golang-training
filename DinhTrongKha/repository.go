package main

import "github.com/jinzhu/gorm"

type Repository struct {
	db *gorm.DB
}

type RepositoryInterface interface {
	DetailUser(user *User) error
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(user *User) error

	DetailPost(post *Post) error
	CreatePost(post *Post) error
	UpdatePost(post *Post) error
}

// CreateUser create new user
func (repo *Repository) CreateUser(user *User) error {
	if db.NewRecord(user) {
		if err := repo.db.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}

// DetailUser get info user
func (repo *Repository) DetailUser(user *User) error {
	result := repo.db.First(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UpdateUser update user
func (repo *Repository) UpdateUser(user *User) error {
	if err := repo.db.Model(&user).Update(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *Repository) DeleteUser(user *User) error {
	if err := repo.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *Repository) DetailPost(post *Post) error {
	result := repo.db.First(&post)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (repo *Repository) CreatePost(post *Post) error {
	if db.NewRecord(post) {
		if err := repo.db.Create(&post).Error; err != nil {
			return err
		}
	}

	return nil
}
func (repo *Repository) UpdatePost(post *Post) error {
	if err := repo.db.Model(&post).Update(&post).Error; err != nil {
		return err
	}

	return nil
}
