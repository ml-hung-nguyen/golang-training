package user

import (
	"encoding/json"
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	Repository RepositoryInterface
}

type UseCaseInterface interface {
	CreateUser(ur *CreateUserRequest) (UserResponse, int, error)
	GetUser(id int) (UserResponse, int, error)
	LoginUser(ur *LoginUserRequest) (JwtToken, int, error)
	UpdateUser(user *User, updatedUser *User) (UserResponse, int, error)
}

func (uc *UseCase) CreateUser(ur *CreateUserRequest) (UserResponse, int, error) {
	response := UserResponse{}
	password, err := bcrypt.GenerateFromPassword([]byte(ur.Password), 14)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}
	var user User
	user.Username = ur.Username
	user.Password = string(password)
	user.FullName = ur.FullName

	err = uc.Repository.CreateUser(&user)
	if err != nil {
		return response, http.StatusUnprocessableEntity, err
	}

	data, err := json.Marshal(&user)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}
	return response, http.StatusOK, nil
}

func (uc *UseCase) GetUser(id int) (UserResponse, int, error) {
	var response UserResponse
	user, err := uc.Repository.FindUser(map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return response, http.StatusUnprocessableEntity, err
		} else {
			return response, http.StatusInternalServerError, err
		}
	}
	data, err := json.Marshal(&user)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}
	return response, http.StatusOK, nil
}

func (uc *UseCase) LoginUser(ur *LoginUserRequest) (JwtToken, int, error) {
	var jwtToken JwtToken
	user, err := uc.Repository.FindUser(map[string]interface{}{"username": ur.Username})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return jwtToken, http.StatusUnprocessableEntity, err
		} else {
			return jwtToken, http.StatusInternalServerError, err
		}
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ur.Password)); err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.Id,
			"username": user.Username,
			"fullname": user.FullName,
		})
		tokenString, err := token.SignedString([]byte("somesecretcode"))
		if err != nil {
			return jwtToken, http.StatusInternalServerError, err
		}
		jwtToken.Token = tokenString
		return jwtToken, http.StatusOK, err
	} else {
		return jwtToken, http.StatusUnauthorized, errors.New("Unauthorize")
	}
}

func (uc *UseCase) UpdateUser(user *User, updatedUser *User) (UserResponse, int, error) {
	var response UserResponse
	err := uc.Repository.UpdateUser(user, updatedUser)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}
	data, err := json.Marshal(&user)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, http.StatusInternalServerError, err
	}
	return response, http.StatusOK, nil
}

func NewUseCase(r *Repository) *UseCase {
	return &UseCase{
		Repository: r,
	}
}
