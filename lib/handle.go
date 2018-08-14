package lib

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
	DB *sql.DB
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = user.Get(userId, h.DB)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(err)
		return
	}
	fmt.Println(user)
	json.NewEncoder(w).Encode(user)
	return
}
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
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

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(err.Error())
		return
	}
	err = user.Update(userId, h.DB)
	if err != nil {
		w.WriteHeader(400)
		fmt.Println(err.Error())
		return
	}
}
