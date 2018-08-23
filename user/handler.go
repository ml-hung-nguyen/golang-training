package user

import (
	"encoding/json"
	. "golang-training/common"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	UseCase UseCaseInterface
}

func NewHandler(db *gorm.DB) *Handler {
	repo := NewRepository(db)
	uc := NewUseCase(repo)
	return &Handler{
		UseCase: uc,
	}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	request := CreateUserRequest{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		JsonResponse(w, http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}
	if len(body) < 1 {
		JsonResponse(w, http.StatusBadRequest, ErrorResponse{Message: "No body"})
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(&request)
	if err != nil {
		JsonResponse(w, http.StatusUnprocessableEntity, ErrorResponse{Message: err.Error()})
		return
	}

	response, status, err := h.UseCase.CreateUser(&request)
	if err != nil {
		JsonResponse(w, status, ErrorResponse{Message: err.Error()})
		return
	}
	JsonResponse(w, status, response)
	return
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		JsonResponse(w, http.StatusBadRequest, ErrorResponse{Message: "Invalid ID"})
		return
	}

	response, status, err := h.UseCase.GetUser(id)
	if err != nil {
		JsonResponse(w, status, ErrorResponse{Message: err.Error()})
		return
	}
	JsonResponse(w, status, response)
	return
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var request LoginUserRequest

	err := ParseForm(r, &request)
	if err != nil {
		JsonResponse(w, http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(&request)
	if err != nil {
		JsonResponse(w, http.StatusUnprocessableEntity, ErrorResponse{Message: err.Error()})
		return
	}
	token, status, err := h.UseCase.LoginUser(&request)
	if err != nil {
		JsonResponse(w, status, ErrorResponse{Message: err.Error()})
		return
	}
	JsonResponse(w, status, token)
	return
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser User
	user := r.Context().Value("user").(User)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		JsonResponse(w, http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	if len(body) < 1 {
		JsonResponse(w, http.StatusBadRequest, ErrorResponse{Message: "No body"})
		return
	}

	err = json.Unmarshal(body, &updateUser)
	if err != nil {
		JsonResponse(w, http.StatusBadRequest, ErrorResponse{Message: "Invalid Json"})
		return
	}

	response, status, err := h.UseCase.UpdateUser(&user, &updateUser)
	if err != nil {
		JsonResponse(w, status, ErrorResponse{Message: err.Error()})
	}
	JsonResponse(w, status, response)
	return
}
