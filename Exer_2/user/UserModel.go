package user

import (
	"time"
)

type User struct {
	Id        int    `form:"-"`
	Username  string `gorm:"username" form:"username"`
	FullName  string `gorm:"full_name" form:"-"`
	Password  string `gorm:"password" form:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
