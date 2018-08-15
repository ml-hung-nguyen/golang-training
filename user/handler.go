package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"github.com/go-playground/form"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	DB   *gorm.DB
	User UserInterface
}

func (u *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	user := User{}
	request := UserRequest{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if len(body) < 1 {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "No body"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &request)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	user.Username = request.Username
	user.Password = string(password)
	user.FullName = request.FullName

	err = u.User.Create(&user, u.DB)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	response := UserResponse{}
	data, err := json.Marshal(&user)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(response)
	return
}

func (u *UserHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	user := User{}

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid ID"})
		return
	}
	err = user.Find(id, u.DB)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	data, err := json.Marshal(&user)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	response := UserResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(response)
	return
}

func (u *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	var request UserRequest

	err := r.ParseForm()
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	decoder := form.NewDecoder()
	err = decoder.Decode(&request, r.Form)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	find := u.DB.Where("username = ?", request.Username).First(&user)
	if find.RecordNotFound() {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "No record"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(request.Password, user.Password)
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.ID,
			"username": user.Username,
			"fullname": user.FullName,
		})
		tokenString, err := token.SignedString([]byte("somesecretcode"))
		if err != nil {
			json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Unauthorize"})
		return
	}
}

func (u *UserHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var updateUser User
	user := r.Context().Value("user").(User)

	err := user.Find(user.ID, u.DB)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(body) < 1 {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "No body"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &updateUser)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "Invalid json"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = u.User.Update(&user, &updateUser, u.DB)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	data, err := json.Marshal(&user)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	response := UserResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(response)
	return
}
