package user

import "time"

type User struct {
	ID        int    `json:id`
	Username  string `json:username`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	FullName  string `json:fullname`
}

type ErrorResponse struct {
	Message string `json:message`
}
