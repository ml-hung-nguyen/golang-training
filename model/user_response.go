package model

type CommonResponse struct {
	Message string `json:"message"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}
