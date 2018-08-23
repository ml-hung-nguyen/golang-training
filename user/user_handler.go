package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/form"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	uc UseCaseInterface
}

func (h *Handler) LoginlUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, CommonResponse{Message: "Internal server response"})
		return
	}

	var userRequest LoginUserRequest
	err = form.NewDecoder().Decode(&userRequest, r.Form)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, CommonResponse{Message: "Internal server response"})
		return
	}

	token, err := h.uc.AuthenUser(userRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ResponseJSON(w, http.StatusNotFound, CommonResponse{Message: "Internal server response"})
		} else if err.Error() == "Unauthorize" {
			ResponseJSON(w, http.StatusUnauthorized, CommonResponse{Message: "Internal server response"})
		} else {
			ResponseJSON(w, http.StatusInternalServerError, CommonResponse{Message: "Internal server response"})
		}
	} else {
		ResponseJSON(w, http.StatusOK, token)
	}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	request := CreateUserRequest{}
	err := ParseForm(r, &request)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, CommonResponse{Message: "Internal server response"})
		return
	}
	user, err := h.uc.CreateUser(&request)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, CommonResponse{Message: "Internal server response"})
		return
	}
	userR := CreateUserResponse{}
	err = ParseJson(&user, &userR)
	ResponseJSON(w, http.StatusOK, user)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, CommonResponse{Message: "Internal server response"})
		return
	}
	user, err := h.uc.GetUser(userId)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, CommonResponse{Message: "Internal server response"})
		return
	}
	userR := GetUserResponse{}
	err = ParseJson(&user, &userR)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, CommonResponse{Message: "Internal server response"})
		return
	}
	ResponseJSON(w, http.StatusOK, userR)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	user, err := h.uc.GetUser(userId)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, CommonResponse{Message: "Internal server response"})
		return
	}
	request := UpdateUserRequest{}
	err = ParseForm(r, &request)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, CommonResponse{Message: "Internal server response"})
		return
	}
	user, err = h.uc.UpdateUser(&user, &request)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, CommonResponse{Message: "Internal server response"})
		return
	}
	userR := UpdateUserResponse{}
	err = ParseJson(&user, &userR)
	ResponseJSON(w, http.StatusOK, user)
}

func ParseJson(user interface{}, response interface{}) error {
	data, err := json.Marshal(&user)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return err
	}
	return nil
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

func NewHandler(uc *UseCase) *Handler {
	return &Handler{
		uc: uc,
	}
}
