package user

type CreateUserRequest struct {
	Username string `form:"username"`
	Fullname string `form:"fullname"`
	Password string `form:"password"`
}

type UpdateUserRequest struct {
	Username string `form:"username"`
	Fullname string `form:"fullname"`
	Password string `form:"password"`
}

type LoginUserRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
