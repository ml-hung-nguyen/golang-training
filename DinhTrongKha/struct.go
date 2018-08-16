package main

import (
	"time"
)

type User struct {
	Id        int    `gorm:"id" form:"-"`
	Username  string `gorm:"username" form:"username"`
	FullName  string `gorm:"full_name" form:"fullname"`
	Password  string `gorm:"password" form:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserResponse struct {
	Id       int    `json:"Id"`
	Username string `json:"Username"`
	FullName string `json:"FullName"`
	password string `json:"-"`
}

type Post struct {
	Id        int    `gorm:"id" form:"-"`
	IdUser    int    `gorm:"id_user" form:"iduser"`
	Content   string `gorm:"content" form:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type PostResponse struct {
	Id      int    `json:"Id"`
	IdUser  int    `json:"IdUser"`
	Content string `json:"Content"`
}
