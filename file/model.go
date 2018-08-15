package user

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type UserInterface interface {
	Find(id int, db *sql.DB) error
	Create(db *sql.DB) error
	Update(id int, db *sql.DB) error
}

type User struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	FullName  string     `json:"full_name"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (h *User) Create(db *sql.DB) error {
	//fmt.Println("Begin Create")
	inserted := 0
	stmt := "insert into users(username, created_at, updated_at, full_name, password) values ($1, $2, $3, $4, $5) returning id"
	err := db.QueryRow(stmt, h.Username, h.CreatedAt, h.UpdatedAt, h.FullName, h.Password).Scan(&inserted)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("End Create")
	return nil
}

func (h *User) Find(id int, db *sql.DB) error {
	stmt := "Select id, username, full_name From users Where Id = $1"
	rows, err := db.Query(stmt, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&h.Id, &h.Username, &h.FullName)
		if err != nil {
			// fmt.Println(err)
			return err
		}
	} else {
		return errors.New("record not found")
	}
	fmt.Println(h)

	return nil
}

func (h *User) Update(id int, db *sql.DB) error {
	statement := "update users set username=$2, full_name=$3, password=$4 where id=$1 returning id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	//fmt.Println(statement)
	defer stmt.Close()
	// randomnumber := 0
	err = stmt.QueryRow(id, h.Username, h.FullName, h.Password).Scan(&h.Id)
	h.Find(id, db)
	return err
}
