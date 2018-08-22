package user

import "time"

type User struct {
	Id        int    `gorm:"id" form:"-"`
	Username  string `gorm:"username" form:"username"`
	FullName  string `gorm:"fullname" form:"fullname"`
	Password  string `gorm:"password" form:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
