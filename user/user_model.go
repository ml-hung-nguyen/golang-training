package user

import "time"

type User struct {
	Id        int    `gorm:"column:id;not null;primary_key"`
	Username  string `gorm:"username"`
	FullName  string `gorm:"fullname"`
	Password  string `gorm:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
