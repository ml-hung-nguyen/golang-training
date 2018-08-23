package model

type MessageResponse struct {
	Message string `json:"message"`
	Errors  string `json:"errors"`
}
