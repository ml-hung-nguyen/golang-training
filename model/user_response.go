package model

type UserResponse struct {
	Id       int    `json:"Id"`
	Username string `json:"Username"`
	FullName string `json:"FullName"`
	password string `json:"-"`
}

type PostResponse struct {
	Id      int    `json:"Id"`
	IdUser  int    `json:"IdUser"`
	Content string `json:"Content"`
}

type CommonResponse struct {
	Message string `json:"message"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}
