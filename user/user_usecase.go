package user

import (
	"encoding/json"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	repo RepositoryInterface
}
type UseCaseInterface interface {
	CreateUser(ur *CreateUserRequest) (UserResponse, error)
	FindUser(id int) (UserResponse, error)
	UpdateUser(user *User, updatedUser *UpdateUserRequest) (UserResponse, error)
	LoginUser(ur *LoginRequest) (TokenResponse, error)
}

func (uc *UseCase) LoginUser(ur *LoginRequest) (TokenResponse, error) {
	var tokenRepo TokenResponse
	user, err := uc.repo.FindUser(map[string]interface{}{"username": ur.Username})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return tokenRepo, err
		} else {
			return tokenRepo, err
		}
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ur.Password)); err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.ID,
			"username": user.Username,
			"fullname": user.FullName,
		})
		tokenString, err := token.SignedString([]byte("somesecretcode"))
		if err != nil {
			return tokenRepo, err
		}
		tokenRepo.Token = tokenString
		return tokenRepo, err
	} else {
		return tokenRepo, errors.New("Unauthorize")
	}
}
func (uc *UseCase) CreateUser(ur *CreateUserRequest) (UserResponse, error) {
	response := UserResponse{}
	password, err := bcrypt.GenerateFromPassword([]byte(ur.Password), 14)
	if err != nil {
		return response, err
	}
	var user User
	user.Username = ur.Username
	user.Password = string(password)
	user.FullName = ur.FullName

	err = uc.repo.CreateUser(&user)
	if err != nil {
		return response, err
	}
	data, err := json.Marshal(&user)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (uc *UseCase) FindUser(id int) (UserResponse, error) {
	var response UserResponse
	user, err := uc.repo.FindUser(map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return response, err
		} else {
			return response, err
		}
	}
	data, err := json.Marshal(&user)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (uc *UseCase) UpdateUser(user *User, updatedUser *UpdateUserRequest) (UserResponse, error) {
	var response UserResponse
	password, err := bcrypt.GenerateFromPassword([]byte(updatedUser.Password), 14)
	if err != nil {
		return response, err
	}
	updatedUser.Password = string(password)
	err = uc.repo.UpdateUser(user, updatedUser)
	if err != nil {
		return response, err
	}
	data, err := json.Marshal(&user)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func NewUseCase(r *Repository) *UseCase {
	return &UseCase{
		repo: r,
	}
}
