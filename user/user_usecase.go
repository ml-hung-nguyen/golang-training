package user

import (
	"golang-training/helper"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	Repository RepositoryInterface
}
type UseCaseInterface interface {
	GetUser(id int) (UserResponse, error)
	CreateUser(ur *CreateUserRequest) (UserResponse, error)
	Login(ur *LoginRequest) (JwtToken, error)
	UpdateUser(ur *UpdateUserRequest) (User, error)
}

func NewUseCase(r *Repository) *UseCase {
	return &UseCase{
		Repository: r,
	}
}

func (uc *UseCase) GetUser(id int) (UserResponse, error) {
	var response UserResponse
	user, err := uc.Repository.FindUser(map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return response, err
		} else {
			return response, err
		}
	}
	err = helper.TranDataJson(user, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (uc *UseCase) CreateUser(ur *CreateUserRequest) (UserResponse, error) {
	log.Println(ur)
	response := UserResponse{}
	password, err := bcrypt.GenerateFromPassword([]byte(ur.Password), 14)
	if err != nil {
		return response, err
	}
	var user User
	user.Username = ur.Username
	user.Password = string(password)
	user.FullName = ur.FullName

	err = uc.Repository.CreateUser(&user)
	if err != nil {
		return response, err
	}
	err = helper.TranDataJson(user, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (uc *UseCase) Login(ur *LoginRequest) (JwtToken, error) {
	var jwtToken JwtToken
	user, err := uc.Repository.FindUser(map[string]interface{}{"username": ur.Username})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return jwtToken, err
		} else {
			return jwtToken, err
		}
	}
	// fmt.Println(user)
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ur.Password)); err == nil {
		// fmt.Println("aaa")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.Id,
			"username": user.Username,
			"fullname": user.FullName,
		})
		tokenString, err := token.SignedString([]byte("secretcode"))
		if err != nil {
			return jwtToken, err
		}
		jwtToken.Token = tokenString
		return jwtToken, nil
	} else {
		return jwtToken, err
	}
}

func (uc *UseCase) UpdateUser(ur *UpdateUserRequest) (User, error) {
	user := User{}
	user.Id = ur.Id
	user.Username = ur.Username
	user.FullName = ur.FullName
	if ur.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(ur.Password), 14)
		if err != nil {
			return user, err
		}
		user.Password = string(password)
	}
	err := uc.Repository.UpdateUser(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
