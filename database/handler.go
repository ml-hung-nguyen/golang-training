package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Handle struct {
	DB *sql.DB
}

func (h *Handle) ShowUser(w http.ResponseWriter, r *http.Request) {
	var user User
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = user.Show(userId, h.DB)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(user)
	return
}
func (h *Handle) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &user)
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode(err)
		return
	}
	err = user.Update(userId, h.DB)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(user)
	return
}
func (h *Handle) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	if len(body) < 1 {
		fmt.Println("No body")
		w.WriteHeader(400)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(400)
		return
	}
	err = user.Create(h.DB)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(400)
		return
	}
}
