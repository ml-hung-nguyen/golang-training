package user

import (
	"golang-training/helper"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/form"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	UseCase UseCaseInterface
}

func NewHandler(db *gorm.DB) *Handler {
	ur := NewRepository(db)
	uc := NewUseCase(ur)
	return &Handler{
		UseCase: uc,
	}
}

func (h *Handler) DetailUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	user, err := h.UseCase.GetUser(id)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		return
	}
	response := UserResponse{}
	err = helper.TranDataJson(user, &response)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, response)
	return
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	request := CreateUserRequest{}
	// err := ParseForm(r, &request)
	err := r.ParseForm()
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	err = form.NewDecoder().Decode(&request, r.Form)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	// log.Println(request)
	validate := validator.New()
	err = validate.Struct(&request)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
		return
	}

	user, err := h.UseCase.CreateUser(&request)
	// log.Println(user)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	response := UserResponse{}
	err = helper.TranDataJson(user, &response)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, response)
	return

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	err := r.ParseForm()
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	err = form.NewDecoder().Decode(&request, r.Form)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(&request)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusUnprocessableEntity, map[string]string{"message": err.Error()})
		return
	}
	token, err := h.UseCase.Login(&request)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	} else {
		helper.RespondWithJSON(w, http.StatusOK, token)
		return
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// var updateUser User
	user := r.Context().Value("user").(User)
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}
	if user.Id != id {
		helper.RespondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
		return
	}
	err = r.ParseForm()
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	request := UpdateUserRequest{}
	request.Id = id
	err = form.NewDecoder().Decode(&request, r.Form)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	user, err = h.UseCase.UpdateUser(&request)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.RespondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
		} else {
			helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
	}
	response := UserResponse{}
	err = helper.TranDataJson(user, &response)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	helper.RespondWithJSON(w, http.StatusOK, response)
	return
}
