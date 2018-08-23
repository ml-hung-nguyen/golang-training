package user

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	Repository RepositoryInterface
}

type UseCaseInterface interface {
	CreateUser(user *CreateUserRequest) (User, error)
	GetUser(int) (User, error)
	UpdateUser(user *User, updateuser *UpdateUserRequest) (User, error)
	AuthenUser(userRequest LoginUserRequest) (string, error)
}

func (uc *UseCase) AuthenUser(userRequest LoginUserRequest) (string, error) {
	var err error
	var user User

	user, err = uc.Repository.FindUser(map[string]interface{}{"username": userRequest.Username})
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.Id,
			"username": user.Username,
			"fullname": user.Fullname,
		})
		tokenString, err := token.SignedString([]byte("secretcode"))
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
	return "", errors.New("Unauthorize")
}

func (uc *UseCase) CreateUser(ur *CreateUserRequest) (User, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(ur.Password), 10)
	var user User
	user.Username = ur.Username
	user.Fullname = ur.Fullname
	user.Password = string(password)
	if err != nil {
		return user, err
	}
	err = uc.Repository.CreateUser(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (uc *UseCase) GetUser(id int) (User, error) {
	var user = User{}
	user, err := uc.Repository.GetUser(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, err
		} else {
			return user, err
		}
	}
	return user, nil
}

func (uc *UseCase) UpdateUser(user *User, updatedUser *UpdateUserRequest) (User, error) {
	err := uc.Repository.UpdateUser(user, updatedUser)
	if err != nil {
		return *user, err
	}
	return *user, nil
}

func NewUseCase(r *Repository) *UseCase {
	return &UseCase{
		Repository: r,
	}
}
