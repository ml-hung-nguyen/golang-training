package file

import (
	"time"
)

type User struct {
	ID        int `gorm:"column:id;not null;primary_key"`
	Username  string
	FullName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
type CreateUserRequest struct {
	Username string `form:"username"`
	FullName string `form:"full_name"`
	Password string `form:"password"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}
type CreateUserResponse struct {
	Token string `json:"token"`
}
type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	FullName string `json:"fullname"`
}
