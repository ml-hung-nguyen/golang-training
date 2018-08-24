package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/form"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	usecase UseCaseInterface
}

func NewHandler(db *gorm.DB) *Handler {
	repo := NewRepository(db)
	uc := NewUseCase(repo)
	return &Handler{
		usecase: uc,
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

func (h *Handler) ShowUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid ID"})
		return
	}
	response, err := h.usecase.FindUser(userId)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	ResponseJSON(w, http.StatusOK, response)
	return
}
func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest

	err := ParseForm(r, &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(&request)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	token, err := h.usecase.LoginUser(&request)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	ResponseJSON(w, http.StatusOK, token)
	return
}
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(User)
	request := UpdateUserRequest{}
	err := ParseForm(r, &request)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, ErrorResponse{Message: "Internal server response"})
		return
	}
	response, err := h.usecase.UpdateUser(&user, &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	ResponseJSON(w, http.StatusOK, response)
	return
}
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := CreateUserRequest{}
	err := ParseForm(r, &request)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, ErrorResponse{Message: "Internal server response"})
		return
	}
	user, err := h.usecase.CreateUser(&request)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, ErrorResponse{Message: "Internal server response"})
		return
	}
	ResponseJSON(w, http.StatusOK, user)
	return
}
func ResponseJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json, _ := json.Marshal(data)
		_, _ = w.Write(json)
	}
}
