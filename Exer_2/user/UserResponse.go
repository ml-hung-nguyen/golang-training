package user

type UserResponse struct {
	Id       int    `json:"Id"`
	Username string `json:"Username"`
	FullName string `json:"FullName"`
	password string `json:"-"`
}
