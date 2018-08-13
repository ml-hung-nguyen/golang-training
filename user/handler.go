package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (u *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		w.WriteHeader(400)
		return
	}
	if len(body) < 1 {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "No body"})
		w.WriteHeader(400)
		return
	}

	err = json.Unmarshal(body, &user)
	fmt.Println(user)
	if u.DB.NewRecord(user) {
		u.DB.Create(&user)
		fmt.Println("Created")
	}
	json.NewEncoder(w).Encode(user)
	return
}

func (u *UserHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	id := chi.URLParam(r, "id")
	_, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid ID"})
		return
	}

	if u.DB.Where("id = ?", id).First(&user).RecordNotFound() {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Record not found"})
		return
	}
	json.NewEncoder(w).Encode(user)
	return
}

func (u *UserHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var user, updateUser User
	id := chi.URLParam(r, "id")
	_, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid ID"})
		return
	}

	if u.DB.Where("id = ?", id).First(&user).RecordNotFound() {
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Record not found"})
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		w.WriteHeader(400)
		return
	}

	if len(body) < 1 {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "No body"})
		w.WriteHeader(400)
		return
	}

	err = json.Unmarshal(body, &updateUser)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid json"})
		w.WriteHeader(400)
		return
	}
	u.DB.Model(&user).Updates(&updateUser)

	json.NewEncoder(w).Encode(user)
	return
}
