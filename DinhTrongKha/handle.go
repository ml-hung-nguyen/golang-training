package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/form"
)

type Handle struct {
	repo RepositoryInterface
}

func (h *Handle) DetailUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id_user"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	user := User{}
	user.Id = id

	err = h.repo.DetailUser(&user)
	if err != nil {
		if err.Error() == "record not found" {
			respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		} else {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	} else {
		response := UserResponse{}
		err = tranDataJson(user, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *Handle) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	user := User{}

	err = form.NewDecoder().Decode(&user, r.Form)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	err = h.repo.CreateUser(&user)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	} else {
		response := UserResponse{}
		err = tranDataJson(user, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *Handle) UpdateUser(w http.ResponseWriter, r *http.Request) {
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

	err = h.repo.DetailUser(&user)
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

	err = h.repo.UpdateUser(&user)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	} else {
		response := UserResponse{}
		err = tranDataJson(user, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *Handle) DeleteUser(w http.ResponseWriter, r *http.Request) {
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

	err = h.repo.DetailUser(&user)
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

	err = h.repo.DeleteUser(&user)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	} else {
		response := UserResponse{}
		err = tranDataJson(user, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *Handle) DetailPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id_post"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	post := Post{}
	post.Id = id

	err = h.repo.DetailPost(&post)
	if err != nil {
		if err.Error() == "record not found" {
			respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		} else {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	} else {
		response := PostResponse{}
		err = tranDataJson(post, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *Handle) CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	post := Post{}

	err = form.NewDecoder().Decode(&post, r.Form)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	err = h.repo.CreatePost(&post)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	} else {
		response := PostResponse{}
		err = tranDataJson(post, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *Handle) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id_post"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	err = r.ParseForm()
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	post := Post{}
	post.Id = id

	err = h.repo.DetailPost(&post)
	if err != nil {
		if err.Error() == "record not found" {
			respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		} else {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return
	}

	err = form.NewDecoder().Decode(&post, r.Form)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	err = h.repo.UpdatePost(&post)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	} else {
		response := PostResponse{}
		err = tranDataJson(post, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}
