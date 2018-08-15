package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/form"
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

	user := User{}
	user.Id = id

	err = h.user.Detail(&user, h.db)
	if err != nil {
		if err.Error() == "record not found" {
			respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		} else {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	} else {
		response := UserResponse{}
		data, err := json.Marshal(&user)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		err = json.Unmarshal(data, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *UserHandle) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	user := User{}

	err = form.NewDecoder().Decode(&user, r.Form)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	err = h.user.Create(&user, h.db)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	} else {
		response := UserResponse{}
		data, err := json.Marshal(&user)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		err = json.Unmarshal(data, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *UserHandle) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id_user"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	err = r.ParseForm()
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	user := User{}
	user.Id = id

	err = h.user.Detail(&user, h.db)
	if err != nil {
		if err.Error() == "record not found" {
			respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		} else {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return
	}

	err = form.NewDecoder().Decode(&user, r.Form)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	err = h.user.Update(&user, h.db)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	} else {
		response := UserResponse{}
		data, err := json.Marshal(&user)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		err = json.Unmarshal(data, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}
