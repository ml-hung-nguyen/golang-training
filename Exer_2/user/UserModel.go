package user

import (
	"time"
)

type User struct {
	Id        int    `gorm:"id" form:"-"`
	Username  string `gorm:"username" form:"username"`
	FullName  string `gorm:"full_name" form:"-"`
	Password  string `gorm:"password" form:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type UserCreateRequest struct {
	Username string `form:"username"`
	FullName string `form:"fullname"`
	Password string `form:"password"`
}

type UserUpdateRequest struct {
	Id       int    `form:"-"`
	Username string `form:"username"`
	FullName string `form:"fullname"`
	Password string `form:"password"`
}

type UserResponse struct {
	Id       int    `json:"Id"`
	Username string `json:"Username"`
	FullName string `json:"FullName"`
	password string `json:"-"`
}
