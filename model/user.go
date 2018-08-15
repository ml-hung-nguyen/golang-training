package model

import (
	"time"
)

type User struct {
	Id        int `gorm:"column:id;not null;primary_key"`
	Username  string
	FullName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
