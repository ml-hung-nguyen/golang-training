package user

import (
	"time"
)

type User struct {
	ID        int `gorm:"column:id;not null;primary_key"`
	Username  string
	FullName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
