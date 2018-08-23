package user

type LoginUserRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	FullName string `json:"fullname"`
}
