package user

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	FullName string `json:"fullname"`
}

type JwtToken struct {
	Token string `json:"token"`
}
