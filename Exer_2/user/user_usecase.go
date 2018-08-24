package user

import (
	"errors"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepo UserRepositoryInterface
}

type UserUseCaseInterface interface {
	AuthenUser(userRequest UserLoginRequest) (string, int, error)
	DetailUserUC(id int) (User, int, error)
	CreateUserUC(userRequest UserCreateRequest) (User, int, error)
	UpdateUserUC(userRequest *UserUpdateRequest) (User, int, error)
	DeleteUserUC(id int) (User, int, error)
}

func NewUserUC(repo *UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepo: repo,
	}
}

func (uc *UserUseCase) AuthenUser(userRequest UserLoginRequest) (string, int, error) {
	var err error
	var user User

	err = uc.UserRepo.DetailUser(&user, map[string]interface{}{"username": userRequest.Username})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", http.StatusNotFound, err
		}
		return "", http.StatusInternalServerError, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.Id,
			"username": user.Username,
			"fullname": user.FullName,
		})
		tokenString, err := token.SignedString([]byte("secretcode"))
		if err != nil {
			return "", http.StatusInternalServerError, err
		}
		return tokenString, http.StatusOK, nil
	}
	return "", http.StatusInternalServerError, errors.New("Unauthorize")
}

func (uc *UserUseCase) DetailUserUC(id int) (User, int, error) {
	var err error
	var status int
	user := User{}
	user.Id = id
	err = uc.UserRepo.DetailUser(&user, nil)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
		} else {
			status = http.StatusInternalServerError
		}
	} else {
		status = http.StatusOK
	}
	return user, status, err
}

func (uc *UserUseCase) CreateUserUC(userRequest UserCreateRequest) (User, int, error) {
	status := http.StatusInternalServerError
	user := User{}
	user.Username = userRequest.Username
	user.FullName = userRequest.FullName
	password, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		return user, status, err
	}
	user.Password = string(password)

	err = uc.UserRepo.CreateUser(&user)
	if err != nil {
		status = http.StatusBadRequest
	} else {
		status = http.StatusOK
	}
	return user, status, err
}

func (uc *UserUseCase) UpdateUserUC(userRequest *UserUpdateRequest) (User, int, error) {
	user := User{}
	status := http.StatusInternalServerError
	user.Id = userRequest.Id
	err := uc.UserRepo.DetailUser(&user, nil)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
			return user, status, err
		}
		return user, status, err
	}

	user.Username = userRequest.Username
	user.FullName = userRequest.FullName
	if userRequest.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
		if err != nil {
			return user, status, err
		}
		user.Password = string(password)
	}

	err = uc.UserRepo.UpdateUser(&user)
	if err == nil {
		status = http.StatusOK
	}
	return user, status, err
}

func (uc *UserUseCase) DeleteUserUC(id int) (User, int, error) {
	user := User{}
	status := http.StatusInternalServerError
	user.Id = id
	err := uc.UserRepo.DetailUser(&user, nil)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			status = http.StatusNotFound
			return user, status, err
		}
		return user, status, err
	}
	err = uc.UserRepo.DeleteUser(&user)
	if err == nil {
		status = http.StatusOK
	}
	return user, status, err
}
