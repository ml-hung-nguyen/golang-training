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
	DB   *gorm.DB
	User UserInterface
}

func (u *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := u.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(body) < 1 {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "No body"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	err = user.Create(u.DB)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(user)
	return
}

func (u *UserHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	user := u.User
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid ID"})
		return
	}
	err = user.Find(id, u.DB)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(user)
	return
}

func (u *UserHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var updateUser User
	user := u.User
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid ID"})
		return
	}

	err = user.Find(id, u.DB)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(body) < 1 {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "No body"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &updateUser)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid json"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = user.Update(&updateUser, u.DB)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(user)
	return
}
