package model

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

type Post struct {
	Id        int    `gorm:"id" form:"-"`
	IdUser    string `gorm:"id_user" form:"iduser"`
	Content   string `gorm:"content" form:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
