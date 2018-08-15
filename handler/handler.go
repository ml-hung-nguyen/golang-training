package handler

import (
	"encoding/json"
	"github.com/at-hungnguyen2/golang-training/model"
	"github.com/at-hungnguyen2/golang-training/repository"
	"github.com/go-playground/form"
	"net/http"
)

type Handler struct {
	repo repository.RepositoryInterface
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := model.CreateUserRequest{}
	err := ParseForm(r, &request)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, model.CommonResponse{Message: "Internal server response"})
		return
	}
	_, err = h.repo.CreateUser(request.Username, request.FullName)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, model.CommonResponse{Message: "Internal server response"})
		return
	}
	ResponseJSON(w, http.StatusOK, model.CreateUserResponse{Token: "token"})
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

func ResponseJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json, _ := json.Marshal(data)
		_, _ = w.Write(json)
	}
}

func NewHandler(repo repository.RepositoryInterface) *Handler {
	return &Handler{
		repo: repo,
	}
}
