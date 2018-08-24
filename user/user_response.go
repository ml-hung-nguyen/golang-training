package user

import "time"

type CommonResponse struct {
	Message string `json:"message"`
}

type CreateUserResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Password string `json:"-"`
}

type UpdateUserResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
