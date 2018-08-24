package user

type CreateUserRequest struct {
	Username string `form:"username"  validate:"required"`
	FullName string `form:"full_name"  validate:"required"`
	Password string `form:"password"`
}
type LoginRequest struct {
	Username string `form:"username"  validate:"required"`
	Password string `form:"password" validate:"required"`
}
type UpdateUserRequest struct {
	Username string `form:"username"`
	FullName string `form:"full_name"`
	Password string `form:"password"`
}
