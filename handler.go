package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	db   *gorm.DB
	User UserInterface
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	user := h.User
	err = json.Unmarshal(body, &user)
	err = user.Create(h.db)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	} else {
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Create successfully"})
	}

}

func (h *Handler) Detail(w http.ResponseWriter, r *http.Request) {
	user := h.User
	userId, err := strconv.Atoi(chi.URLParam(r, "id_user"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	err = user.Detail(userId, h.db)

	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	} else {
		respondwithJSON(w, http.StatusCreated, user)
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	user := h.User
	id, err := strconv.Atoi(chi.URLParam(r, "id_user"))
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
