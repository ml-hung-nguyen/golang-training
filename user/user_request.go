package user

type CreateUserRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname"`
}

type LoginRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Id       int    `json:"Id" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	FullName string `json:"fullname" form:"fullname"`
}
