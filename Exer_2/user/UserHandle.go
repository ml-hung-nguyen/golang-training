package user

import (
	"example/Exer_2/helper"
	"example/Exer_2/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/form"
	"github.com/jinzhu/gorm"
)

type UserHandle struct {
	UserUC UserUseCaseInterface
}

func (h *UserHandle) LoginlUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
		return
	}

	var userRequest User
	err = form.NewDecoder().Decode(&userRequest, r.Form)
	if err != nil {
		helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
		return
	}

	token, err := h.UserUC.AuthenUser(userRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.RespondwithJSON(w, http.StatusNotFound, model.MessageResponse{Message: err.Error()})
		} else if err.Error() == "Unauthorize" {
			helper.RespondwithJSON(w, http.StatusUnauthorized, model.MessageResponse{Message: err.Error()})
		} else {
			helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
		}
	} else {
		helper.RespondwithJSON(w, http.StatusOK, model.TokenResponse{Token: token})
	}
}

func (h *UserHandle) DetailUserHandle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id_user"))
	if err != nil {
		helper.RespondwithJSON(w, http.StatusBadRequest, model.MessageResponse{Message: err.Error()})
		return
	}

	user, err := h.UserUC.DetailUserUC(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.RespondwithJSON(w, http.StatusNotFound, model.MessageResponse{Message: err.Error()})
		} else {
			helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
		}
	} else {
		response := UserResponse{}
		err = helper.TranDataJson(user, &response)
		if err != nil {
			helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
			return
		}
		helper.RespondwithJSON(w, http.StatusOK, response)
	}
}

func (h *UserHandle) CreateUserHandle(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helper.RespondwithJSON(w, http.StatusBadRequest, model.MessageResponse{Message: err.Error()})
		return
	}
	userRequest := UserCreateRequest{}

	err = form.NewDecoder().Decode(&userRequest, r.Form)
	if err != nil {
		helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
		return
	}

	user, err := h.UserUC.CreateUserUC(userRequest)
	if err != nil {
		helper.RespondwithJSON(w, http.StatusBadRequest, model.MessageResponse{Message: err.Error()})
	} else {
		response := UserResponse{}
		err = helper.TranDataJson(user, &response)
		if err != nil {
			helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
			return
		}
		helper.RespondwithJSON(w, http.StatusOK, response)
	}
}

func (h *UserHandle) UpdateUserHandle(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(User)
	id, err := strconv.Atoi(chi.URLParam(r, "id_user"))
	if err != nil {
		helper.RespondwithJSON(w, http.StatusBadRequest, model.MessageResponse{Message: err.Error()})
		return
	}

	if user.Id != id {
		helper.RespondwithJSON(w, http.StatusUnauthorized, model.MessageResponse{Message: "Unauthorize"})
		return
	}

	err = r.ParseForm()
	if err != nil {
		helper.RespondwithJSON(w, http.StatusBadRequest, model.MessageResponse{Message: err.Error()})
		return
	}
	userRequest := UserUpdateRequest{}
	userRequest.Id = id

	err = form.NewDecoder().Decode(&userRequest, r.Form)
	if err != nil {
		helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
		return
	}

	user, err = h.UserUC.UpdateUserUC(&userRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.RespondwithJSON(w, http.StatusNotFound, model.MessageResponse{Message: err.Error()})
		} else {
			helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
		}
	} else {
		response := UserResponse{}
		err = helper.TranDataJson(user, &response)
		if err != nil {
			helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
			return
		}
		helper.RespondwithJSON(w, http.StatusOK, response)
	}
}

func (h *UserHandle) DeleteUserHandle(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id_user"))
	if err != nil {
		helper.RespondwithJSON(w, http.StatusBadRequest, model.MessageResponse{Message: err.Error()})
		return
	}

	user, err := h.UserUC.DeleteUserUC(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helper.RespondwithJSON(w, http.StatusNotFound, model.MessageResponse{Message: err.Error()})
		} else {
			helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
		}
	} else {
		response := UserResponse{}
		err = helper.TranDataJson(user, &response)
		if err != nil {
			helper.RespondwithJSON(w, http.StatusInternalServerError, model.MessageResponse{Message: err.Error()})
			return
		}
		helper.RespondwithJSON(w, http.StatusOK, response)
	}
}
