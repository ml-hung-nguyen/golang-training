package user

import (
	"encoding/json"
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	Repository RepositoryInterface
}

type UseCaseInterface interface {
	CreateUser(ur *CreateUserRequest) (UserResponse, error)
	GetUser(id int) (UserResponse, error)
	LoginUser(ur *LoginUserRequest) (JwtToken, error)
	UpdateUser(updatedUser *User) (UserResponse, error)
}

func (uc *UseCase) CreateUser(ur *CreateUserRequest) (UserResponse, error) {
	reponse := UserResponse{}
	password, err := bcrypt.GenerateFromPassword([]byte(ur.Password), 14)
	if err != nil {
		return reponse, err
	}
	var user User
	user.Username = ur.Username
	user.Password = string(password)
	user.Fullname = ur.Fullname

	err = uc.Repository.CreateUser(&user)
	if err != nil {
		return reponse, err
	}

	data, err := json.Marshal(&user)
	if err != nil {
		return reponse, err
	}
	err = json.Unmarshal(data, &reponse)
	if err != nil {
		return reponse, err
	}
	return reponse, nil
}

func (uc *UseCase) GetUser(id int) (UserResponse, error) {
	var reponse UserResponse
	user, err := uc.Repository.FindUser(map[string]interface{}{"id": id})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return reponse, err
		} else {
			return reponse, err
		}
	}
	data, err := json.Marshal(&user)
	if err != nil {
		return reponse, err
	}
	err = json.Unmarshal(data, &reponse)
	if err != nil {
		return reponse, err
	}
	return reponse, nil
}

func (uc *UseCase) LoginUser(ur *LoginUserRequest) (JwtToken, error) {
	var jwtToken JwtToken
	user, err := uc.Repository.FindUser(map[string]interface{}{"username": ur.Username})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return jwtToken, err
		} else {
			return jwtToken, err
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
			return jwtToken, err
		}
		jwtToken.Token = tokenString
		return jwtToken, err
	} else {
		return jwtToken, errors.New("Unauthorize")
	}
}

func (uc *UseCase) UpdateUser(updatedUser *User) (UserResponse, error) {
	var reponse UserResponse
	user, err := uc.Repository.UpdateUser(updatedUser)
	if err != nil {
		return reponse, err
	}
	data, err := json.Marshal(&user)
	if err != nil {
		return reponse, err
	}

	err = json.Unmarshal(data, &reponse)
	if err != nil {
		return reponse, err
	}
	return reponse, nil
}

func NewUseCase(r *Repository) *UseCase {
	return &UseCase{
		Repository: r,
	}
}
