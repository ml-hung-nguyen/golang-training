package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type UserHandle struct {
	db   *gorm.DB
	user UserInterface
}

func (h *UserHandle) Detail(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id_user"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	user := h.user

	err = user.Detail(id, h.db)
	if err != nil {
		if err.Error() == "record not found" {
			respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		} else {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	} else {
		respondwithJSON(w, http.StatusOK, user)
	}
}

func (h *UserHandle) Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	user := h.user
	err = json.Unmarshal(body, &user)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	err = user.Create(h.db)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	} else {
		respondwithJSON(w, http.StatusOK, map[string]string{"message": "Create successfully"})
	}
}

func (h *UserHandle) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id_user"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	body, err := ioutil.ReadAll(r.Body)
	user := h.user
	var userUpdate User
	err = json.Unmarshal(body, &userUpdate)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	err = user.Detail(id, h.db)
	if err != nil {
		if err.Error() == "record not found" {
			respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		} else {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return
	}
	err = user.Update(userUpdate, h.db)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	} else {
		respondwithJSON(w, http.StatusOK, map[string]string{"message": "Update successfully"})
	}
}
