package user

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepo UserRepositoryInterface
}

type UserUseCaseInterface interface {
	AuthenUser(userRequest User) (string, error)
	DetailUserUC(id int) (User, error)
	CreateUserUC(userRequest UserCreateRequest) (User, error)
	UpdateUserUC(userRequest *UserUpdateRequest) (User, error)
	DeleteUserUC(id int) (User, error)
}

func (uc *UserUseCase) AuthenUser(userRequest User) (string, error) {
	var err error
	var user User

	err = uc.UserRepo.DetailUser(&user, map[string]interface{}{"username": userRequest.Username})
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":       user.Id,
			"username": user.Username,
			"fullname": user.FullName,
		})
		tokenString, err := token.SignedString([]byte("secretcode"))
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
	return "", errors.New("Unauthorize")
}

func (uc *UserUseCase) DetailUserUC(id int) (User, error) {
	var err error
	user := User{}
	user.Id = id
	err = uc.UserRepo.DetailUser(&user, nil)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (uc *UserUseCase) CreateUserUC(userRequest UserCreateRequest) (User, error) {
	user := User{}
	user.Username = userRequest.Username
	user.FullName = userRequest.FullName
	password, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
	if err != nil {
		return user, err
	}
	user.Password = string(password)

	err = uc.UserRepo.CreateUser(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (uc *UserUseCase) UpdateUserUC(userRequest *UserUpdateRequest) (User, error) {
	user := User{}
	user.Id = userRequest.Id
	err := uc.UserRepo.DetailUser(&user, nil)
	if err != nil {
		return user, err
	}

	if userRequest.Username != "" {
		user.Username = userRequest.Username
	}
	if userRequest.FullName != "" {
		user.FullName = userRequest.FullName
	}
	if userRequest.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 10)
		if err != nil {
			return user, err
		}
		user.Password = string(password)
	}

	err = uc.UserRepo.UpdateUser(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (uc *UserUseCase) DeleteUserUC(id int) (User, error) {
	user := User{}
	user.Id = id
	err := uc.UserRepo.DetailUser(&user, nil)
	if err != nil {
		return user, err
	}
	err = uc.UserRepo.DeleteUser(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
