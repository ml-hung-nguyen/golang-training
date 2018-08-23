package user

import (
	"example/Exer_1/golang-training/Exer_2/helper"
	"example/Exer_1/golang-training/Exer_2/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/form"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v9"
)

type UserHandler struct {
	UserUC UserUseCaseInterface
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	userRepo := NewUserRepo(db)
	userUC := NewUserUC(userRepo)
	return &UserHandler{
		UserUC: userUC,
	}
}

func (h *UserHandler) LoginlUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error(), Errors: err})
		return
	}

	var userRequest UserLoginRequest
	err = form.NewDecoder().Decode(&userRequest, r.Form)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error(), Errors: err})
		return
	}

	validate := validator.New()
	err = validate.Struct(&userRequest)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusUnprocessableEntity, model.MessageResponse{Message: err.Error(), Errors: err})
		return
	}

	token, status, err := h.UserUC.AuthenUser(userRequest)
	if err != nil {
		helper.RespondWithJSON(w, status, model.MessageResponse{Message: err.Error(), Errors: err})
	} else {
		helper.RespondWithJSON(w, status, model.TokenResponse{Token: token})
	}
}

func (h *UserHandler) DetailUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id_user"))
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, model.MessageResponse{Message: err.Error(), Errors: err})
		return
	}

	user, status, err := h.UserUC.DetailUserUC(id)
	if err != nil {
		helper.RespondWithJSON(w, status, model.MessageResponse{Message: err.Error(), Errors: err})
	} else {
		response := UserResponse{}
		err = helper.TranDataJson(user, &response)
		if err != nil {
			helper.RespondWithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error(), Errors: err})
			return
		}
		helper.RespondWithJSON(w, status, response)
	}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, model.MessageResponse{Message: err.Error(), Errors: err})
		return
	}
	userRequest := UserCreateRequest{}

	err = form.NewDecoder().Decode(&userRequest, r.Form)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error(), Errors: err})
		return
	}

	validate := validator.New()
	err = validate.Struct(&userRequest)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusUnprocessableEntity, model.MessageResponse{Message: err.Error(), Errors: err})
		return
	}

	user, status, err := h.UserUC.CreateUserUC(userRequest)
	if err != nil {
		helper.RespondWithJSON(w, status, model.MessageResponse{Message: err.Error(), Errors: err})
	} else {
		response := UserResponse{}
		err = helper.TranDataJson(user, &response)
		if err != nil {
			helper.RespondWithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error(), Errors: err})
			return
		}
		helper.RespondWithJSON(w, status, response)
	}
}

func (h *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(User)
	id, err := strconv.Atoi(chi.URLParam(r, "id_user"))
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, model.MessageResponse{Message: err.Error(), Errors: err})
		return
	}

	if user.Id != id {
		helper.RespondWithJSON(w, http.StatusUnauthorized, model.MessageResponse{Message: "Unauthorize"})
		return
	}

	err = r.ParseForm()
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, model.MessageResponse{Message: err.Error(), Errors: err})
		return
	}
	userRequest := UserUpdateRequest{}
	userRequest.Id = id

	err = form.NewDecoder().Decode(&userRequest, r.Form)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error(), Errors: err})
		return
	}

	user, status, err := h.UserUC.UpdateUserUC(&userRequest)
	if err != nil {
		helper.RespondWithJSON(w, status, model.MessageResponse{Message: err.Error(), Errors: err})
	} else {
		response := UserResponse{}
		err = helper.TranDataJson(user, &response)
		if err != nil {
			helper.RespondWithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error(), Errors: err})
			return
		}
		helper.RespondWithJSON(w, status, response)
	}
}

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id_user"))
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, model.MessageResponse{Message: err.Error(), Errors: err})
		return
	}

	user, status, err := h.UserUC.DeleteUserUC(id)
	if err != nil {
		helper.RespondWithJSON(w, status, model.MessageResponse{Message: err.Error(), Errors: err})
	} else {
		response := UserResponse{}
		err = helper.TranDataJson(user, &response)
		if err != nil {
			helper.RespondWithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error(), Errors: err})
			return
		}
		helper.RespondWithJSON(w, status, response)
	}
}
