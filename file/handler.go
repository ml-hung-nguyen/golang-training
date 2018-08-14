package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Handler struct {
	DB   *sql.DB
	User UserInterface
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := h.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(body) < 1 {
		w.WriteHeader(400)
		fmt.Println("No body")
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(err.Error())
		return
	}
	err = user.Create(h.DB)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(err.Error())
		return
	}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	user := h.User
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = user.Find(userId, h.DB)
	if err != nil {
		//mt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(user)
	return
}

func (u *User) Update(id int, db *sql.DB) error {
	statement := "Update users set username=$2, full_name=$6 where id=$1 returning id"
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id, u.Username, u.FullName).Scan(&u.Id)
	u.Find(id, db)
	return err
}
