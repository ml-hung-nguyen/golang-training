package user

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
