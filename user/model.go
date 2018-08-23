package user

import (
	"time"
)

type User struct {
	ID        int    `gorm:"column:id;not null;primary_key"`
	Username  string `gorm:"column:username;not null"`
	Password  string `gorm:"column:password;not null"`
	Fullname  string `gorm:"column:fullname;null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
