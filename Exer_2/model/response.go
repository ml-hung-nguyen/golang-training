package model

type MessageResponse struct {
	Message string
	Errors  error
}

type TokenResponse struct {
	Token string
}
