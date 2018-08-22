package user

import (
	"baitapgo_ngay1/golang-training/model"
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
	CreateUser(ur *CreateUserRequest) (User, int, error)
	GetUser(id int) (User, int, error)
	LoginUser(ur *LoginUserRequest) (model.JwtToken, int, error)
	UpdateUser(user *User, updatedUser *User) (User, int, error)
}

func (uc *UseCase) CreateUser(ur *CreateUserRequest) (User, int, error) {
	u := User{}
	password, err := bcrypt.GenerateFromPassword([]byte(ur.Password), 14)
	if err != nil {
		return u, http.StatusInternalServerError, err
	}
	var user User
	user.Username = ur.Username
	user.Password = string(password)
	user.Fullname = ur.Fullname

	err = uc.Repository.CreateUser(&user)
	if err != nil {
		return u, http.StatusUnprocessableEntity, err
	}

	data, err := json.Marshal(&user)
	if err != nil {
		return u, http.StatusInternalServerError, err
	}
	err = json.Unmarshal(data, &u)
	if err != nil {
		return u, http.StatusInternalServerError, err
	}
	return u, http.StatusOK, nil
}

func (uc *UseCase) GetUser(id int) (User, int, error) {
	var reponsitory Repository
	user, err := uc.Repository.FindUser(map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, http.StatusUnprocessableEntity, err
		} else {
			return user, http.StatusInternalServerError, err
		}
	}
	data, err := json.Marshal(&user)
	if err != nil {
		return user, http.StatusInternalServerError, err
	}
	err = json.Unmarshal(data, &reponsitory)
	if err != nil {
		return user, http.StatusInternalServerError, err
	}
	return user, http.StatusOK, nil
}

func (uc *UseCase) LoginUser(ur *LoginUserRequest) (model.JwtToken, int, error) {
	var jwtToken model.JwtToken
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
			"id":       user.ID,
			"username": user.Username,
			"fullname": user.Fullname,
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

func (uc *UseCase) UpdateUser(user *User, updatedUser *User) (User, int, error) {
	var u User
	err := uc.Repository.UpdateUser(user, updatedUser)
	if err != nil {
		return u, http.StatusInternalServerError, err
	}
	data, err := json.Marshal(&user)
	if err != nil {
		return u, http.StatusInternalServerError, err
	}

	err = json.Unmarshal(data, &u)
	if err != nil {
		return u, http.StatusInternalServerError, err
	}
	return u, http.StatusOK, nil
}

func NewUseCase(r *Repository) *UseCase {
	return &UseCase{
		Repository: r,
	}
}
