package file

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/form"
	"github.com/jinzhu/gorm"
)

type Handler struct {
	Repository RepositoryInterface
}

func NewHandler(r *Repository) *Handler {
	return &Handler{
		Repository: r,
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
	user := User{}
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid ID"})
		return
	}
	user, err = h.Repository.FindUser(map[string]interface{}{"id": userId})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
			return
		}
	}
	data, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	response := UserResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(response)
	return
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUser User
	user := User{}
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	body, _ := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	if len(body) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "No body"})
		return
	}
	err = json.Unmarshal(body, &updateUser)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid json"})
		return
	}
	user, err = h.Repository.FindUser(map[string]interface{}{"id": userId})
	if err != nil {
		return
	}
	err = h.Repository.UpdateUser(&user, &updateUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	data, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	response := UserResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(response)
	return
}
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := User{}
	request := CreateUserRequest{}
	err := ParseForm(r, &request)
	user.Username = request.Username
	user.Password = request.Password
	user.FullName = request.FullName
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, ErrorResponse{Message: "Internal server response"})
		return
	}
	err = h.Repository.CreateUser(&user)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, ErrorResponse{Message: "Internal server response"})
		return
	}
	ResponseJSON(w, http.StatusOK, CreateUserResponse{Token: "token"})
}
func ResponseJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json, _ := json.Marshal(data)
		_, _ = w.Write(json)
	}
}
