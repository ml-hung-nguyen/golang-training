package user

import (
	"time"
)

type User struct {
	ID        int
	Username  string
	Password  string
	FullName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
