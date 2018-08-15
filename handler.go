package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type HandlerUser struct {
	db   *gorm.DB
	User UserInterface
}

func (h *HandlerUser) Detail(w http.ResponseWriter, r *http.Request) {
	UserId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	user := h.User
	err = user.Detail(UserId, h.db)
	if err != nil {
		respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
	} else {
		respondwithJSON(w, http.StatusOK, user)
	}
}

func (h *HandlerUser) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	user := h.User
	err = json.Unmarshal(body, &user)
	fmt.Println(user)
	err = user.Create(h.db)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	} else {
		respondwithJSON(w, http.StatusCreated, user)
	}
}

func (h *HandlerUser) Update(w http.ResponseWriter, r *http.Request) {
	user := h.User
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	body, err := ioutil.ReadAll(r.Body)
	var userUpdate User
	err = json.Unmarshal(body, &userUpdate)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	err = user.Detail(id, h.db)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	err = user.Update(userUpdate, h.db)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	} else {
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Update successfully"})
	}
}
