package user

import (
	"baitapgo_ngay1/golang-training/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/form"
	validator "gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	UseCase UseCaseInterface
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

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	request := CreateUserRequest{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
		return
	}
	if len(body) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.MessageResponse{Message: "No body"})
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(&request)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
		return
	}

	response, status, err := h.UseCase.CreateUser(&request)
	w.WriteHeader(status)

	if err != nil {
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(response)
	return
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.MessageResponse{Message: "Invalid ID"})
		return
	}

	response, status, err := h.UseCase.GetUser(id)
	w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(response)
	return
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var request LoginUserRequest

	err := ParseForm(r, &request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(&request)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
		return
	}
	token, status, err := h.UseCase.LoginUser(&request)
	w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(token)
	return
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser User
	//fmt.Println("Handler")
	user, ok := r.Context().Value("user").(User)
	// user = user.(User)
	if !ok {
		fmt.Println(ok)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
		return
	}

	if len(body) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.MessageResponse{Message: "No body"})
		return
	}

	err = json.Unmarshal(body, &updateUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.MessageResponse{Message: "Invalid json"})
		return
	}

	response, status, err := h.UseCase.UpdateUser(&user, &updateUser)
	w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
	}
	json.NewEncoder(w).Encode(response)
	return
}

func NewHandler(uc *UseCase) *Handler {
	return &Handler{
		UseCase: uc,
	}
}
