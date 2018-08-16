package lib

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int
	Username  string
	Fullname  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (u *User) Get(id int, db *sql.DB) error {
	statement := "select id, username, fullname, password from users where id = $1"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&u.Id, &u.Username, &u.Fullname, &u.Password)
	return err

}

func (u *User) Create(db *sql.DB) error {
	statement := "insert into users (username, fullname, password, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.Username, u.Fullname, u.Password, u.CreatedAt, u.UpdatedAt).Scan(&u.Id)
	return err

}

func (u *User) Update(id int, db *sql.DB) error {
	statement := "update users set username=$1, fullname=$2, password=$3 where id=$4 returning id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.Username, u.Fullname, u.Password, id).Scan(&u.Id)
	return err

}
