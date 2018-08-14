package database

import (
	"database/sql"
	"time"
)

type User struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Fullname string    `json:"fullname"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
	DeleteAt time.Time `json:"deleted_at"`
}

func (u *User) Create(db *sql.DB) error {
	statement := "insert into users (username, full_name, created_at, updated_at) values($1,$2,$3,$4) returning ID"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.Username, u.Fullname, u.CreateAt, u.UpdateAt).Scan(&u.ID)
	return err
}
func (u *User) Update(id int, db *sql.DB) error {
	statement := "Update users set username=$2, full_name=$3, created_at=$4, updated_at=$5 where id=$1 returning id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id, u.Username, u.Fullname, u.CreateAt, u.UpdateAt).Scan(&u.ID)
	u.Show(id, db)
	return err
}
func (u *User) Show(id int, db *sql.DB) error {
	stmt := "Select id,username,full_name,created_at,updated_at from users Where id = $1"
	rows, err := db.Query(stmt, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Username, &u.Fullname, &u.CreateAt, &u.UpdateAt)
		if err != nil {
			return err
		}
	}
	return nil
}
