package handler

import (
	"encoding/json"
	"fmt"
	"golang-training/model"
	"golang-training/repository"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-playground/form"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	repo repository.RepositoryInterface
}

func NewHandler(repo repository.RepositoryInterface) *Handler {
	return &Handler{
		repo: repo,
	}
}

func ParseForm(r *http.Request, i interface{}) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	decoder := form.NewDecoder()
	err = decoder.Decode(&i, r.Form)
	return err
}

func respondwithJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json, _ := json.Marshal(data)
		_, _ = w.Write(json)
	}
}

func tranDataJson(origin interface{}, response interface{}) error {
	data, err := json.Marshal(&origin)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) DetailUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// user := h.User
	user, err = h.repo.FindUser(map[string]interface{}{"id": id})
	if err != nil {
		respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
	} else {
		respondwithJSON(w, http.StatusOK, user)
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	user := model.User{}
	err = form.NewDecoder().Decode(&user, r.Form)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	// fmt.Println(password)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	user.Password = string(password)
	err = h.repo.CreateUser(&user)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	} else {
		response := model.UserResponse{}
		err = tranDataJson(user, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	err = r.ParseForm()
	fmt.Println(r.Form)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	user := model.User{}
	user.Id = id
	user, err = h.repo.FindUser(map[string]interface{}{"id": id})
	if err != nil {
		if err.Error() == "record not found" {
			respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
			return
		} else {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		return
	}
	err = form.NewDecoder().Decode(&user, r.Form)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	err = h.repo.UpdateUser(&user)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	} else {
		response := model.UserResponse{}
		err = tranDataJson(user, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	err = r.ParseForm()
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	user := model.User{}
	user.Id = id
	user, err = h.repo.FindUser(map[string]interface{}{"id": id})
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
		response := model.UserResponse{}
		err = tranDataJson(user, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *Handler) LoginlUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	var user, userRequest model.User
	err = form.NewDecoder().Decode(&userRequest, r.Form)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	user, err = h.repo.FindUser(map[string]interface{}{"username": userRequest.Username})
	if err != nil {
		if err.Error() == "record not found" {
			respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		} else {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.Id,
			"username": user.Username,
			"fullname": user.FullName,
		})
		tokenString, err := token.SignedString([]byte("secretcode"))
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, err)
			return
		}
		respondwithJSON(w, http.StatusOK, map[string]string{"Token": tokenString})
	} else {
		respondwithJSON(w, http.StatusUnauthorized, map[string]string{"message": "Unauthorize"})
	}
}

func (h *Handler) DetailPost(w http.ResponseWriter, r *http.Request) {
	post := model.Post{}
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// user := h.User
	post, err = h.repo.DetailPost(map[string]interface{}{"id": id})
	if err != nil {
		respondwithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
	} else {
		respondwithJSON(w, http.StatusOK, post)
	}
}

func (h *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	post := model.Post{}
	fmt.Println(post)
	err = form.NewDecoder().Decode(&post, r.Form)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	err = h.repo.CreatePost(&post)
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	} else {
		response := model.PostResponse{}
		err = tranDataJson(post, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}

func (h *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	err = r.ParseForm()
	if err != nil {
		respondwithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	post := model.Post{}
	post.Id = id
	post, err = h.repo.DetailPost(map[string]interface{}{"id": id})
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
		response := model.PostResponse{}
		err = tranDataJson(post, &response)
		if err != nil {
			respondwithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
			return
		}
		respondwithJSON(w, http.StatusOK, response)
	}
}
