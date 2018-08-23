package user

import (
	"time"
)

type User struct {
	Id        int    `gorm:"id;not null;primary_key"`
	Username  string `gorm:"username;not null;unique"`
	FullName  string `gorm:"full_name"`
	Password  string `gorm:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
