package model

type CreateUserRequest struct {
	Username string `form:"username"`
	FullName string `form:"full_name"`
	Password string `form:"password"`
}
