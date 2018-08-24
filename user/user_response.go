package user

type UserResponse struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Password string `json:"-"`
}

type JwtToken struct {
	Token string `json:"token"`
}
