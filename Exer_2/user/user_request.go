package user

type UserLoginRequest struct {
	Username string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type UserCreateRequest struct {
	Username string `form:"username" validate:"required"`
	FullName string `form:"fullname"`
	Password string `form:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Id       int    `form:"-"`
	Username string `form:"username"`
	FullName string `form:"fullname"`
	Password string `form:"password"`
}
