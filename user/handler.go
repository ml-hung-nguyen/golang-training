package user

import (
	"baitapgo_ngay1/golang-training/common"
	"baitapgo_ngay1/golang-training/model"
	"encoding/json"
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

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	request := CreateUserRequest{}
	err := common.ParseForm(r, &request)
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

	response, err := h.UseCase.CreateUser(&request)
	//w.WriteHeader(status)

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

	response, err := h.UseCase.GetUser(id)
	//w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(model.MessageResponse{Message: "err.Error()"})
		return
	}

	json.NewEncoder(w).Encode(response)
	return
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var request LoginUserRequest

	err := common.ParseForm(r, &request)
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
	token, err := h.UseCase.LoginUser(&request)
	//w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(token)
	return
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser User
	var request UpdateUserRequest
	//fmt.Println("Handler")
	//ok := r.Context().Value("user").(User)
	// user = user.(User)
	// if !ok {
	// 	fmt.Println(ok)
	// 	return
	// }
	err := common.ParseForm(r, &request)
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

	response, err := h.UseCase.UpdateUser(&updateUser)
	//w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(model.MessageResponse{Message: err.Error()})
	}
	json.NewEncoder(w).Encode(response)
	return
}

func NewHandler(db *gorm.DB) *Handler {
	ur := NewRepository(db)
	uc := NewUseCase(ur)
	return &Handler{
		UseCase: uc,
	}
}
