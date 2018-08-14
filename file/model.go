package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

type UserInterface interface {
	Find(id int, db *sql.DB) error
	Create(db *sql.DB) error
	Update(id int, db *sql.DB) error
}

type User struct {
	Id        int        `json:"id"`
	Username  string     `json:"username"`
	FullName  string     `json:"fullname"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (h *User) Create(db *sql.DB) error {
	//fmt.Println("Begin Create")
	inserted := 0
	stmt := "insert into users(username,created_at, updated_at, full_name) values ($1, $2, $3, $4) returning id"
	err := db.QueryRow(stmt, h.Username, h.CreatedAt, h.UpdatedAt, h.FullName).Scan(&inserted)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("End Create")
	return nil
}

func (h *User) Find(id int, db *sql.DB) error {
	stmt := "Select * From users Where Id = $1"
	rows, err := db.Query(stmt, id)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&h.Id, &h.Username, &h.CreatedAt, &h.UpdatedAt, &h.FullName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &user)
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Println(user)
	err = user.Update(userId, h.DB)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	// _ = user.Show(userId, h.DB)
	json.NewEncoder(w).Encode(user)
	return
}
