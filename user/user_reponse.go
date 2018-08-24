package user

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
type TokenResponse struct {
	Token string `json:"token"`
}
