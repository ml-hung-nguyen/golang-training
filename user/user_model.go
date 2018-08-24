package user

import (
	"time"
)

type User struct {
	Id        int `gorm:"column:id;not null;primary_key"`
	Username  string
	Fullname  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
