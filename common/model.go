package common

type ErrorResponse struct {
	Message string `json:"message"`
	Errors  error  `json:"error"`
}
