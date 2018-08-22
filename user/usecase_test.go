package user_test

import (
	"golang-training/user"
	"net/http"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type UseCaseMock struct {
	Status int
	Errors error
}

func (um *UseCaseMock) CreateUser(ur *user.CreateUserRequest) (user.UserResponse, int, error) {
	return user.UserResponse{}, um.Status, um.Errors
}

func (um *UseCaseMock) GetUser(id int) (user.UserResponse, int, error) {
	return user.UserResponse{}, um.Status, um.Errors
}

func (um *UseCaseMock) LoginUser(ur *user.LoginUserRequest) (user.JwtToken, int, error) {
	return user.JwtToken{}, um.Status, um.Errors
}

func (um *UseCaseMock) UpdateUser(u *user.User, updatedUser *user.User) (user.UserResponse, int, error) {
	return user.UserResponse{}, um.Status, um.Errors
}

func TestUseCaseCreateUser(t *testing.T) {
	uc := user.UseCase{
		Repository: &RepositoryMock{},
	}
	request := user.CreateUserRequest{
		Username: "test",
		Password: "test",
	}
	_, status, err := uc.CreateUser(&request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
}

func TestUseCaseCreateUserFail(t *testing.T) {
	uc := user.UseCase{
		Repository: &RepositoryMock{
			Errors: gorm.ErrInvalidSQL,
		},
	}
	request := user.CreateUserRequest{
		Username: "test",
	}
	_, status, err := uc.CreateUser(&request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, status)
}

func TestUseCaseGetUser(t *testing.T) {
	uc := user.UseCase{
		Repository: &RepositoryMock{
			Errors: nil,
		},
	}
	_, status, err := uc.GetUser(1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
}

func TestUseCaseGetUserFail(t *testing.T) {
	uc := user.UseCase{
		Repository: &RepositoryMock{
			Errors: gorm.ErrRecordNotFound,
		},
	}
	_, status, err := uc.GetUser(1)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, status)

	uc = user.UseCase{
		Repository: &RepositoryMock{
			Errors: gorm.ErrInvalidSQL,
		},
	}
	_, status, err = uc.GetUser(1)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestUseCaseLoginUser(t *testing.T) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("abc"), 14)
	uc := user.UseCase{
		Repository: &RepositoryMock{
			Errors:     nil,
			Conditions: map[string]interface{}{"password": string(hash)},
		},
	}
	request := user.LoginUserRequest{Username: "test", Password: "abc"}
	_, status, err := uc.LoginUser(&request)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, status)
}

func TestUseCaseLoginUserFailNoRecord(t *testing.T) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("abc"), 14)
	uc := user.UseCase{
		Repository: &RepositoryMock{
			Errors:     gorm.ErrRecordNotFound,
			Conditions: map[string]interface{}{"password": string(hash)},
		},
	}
	request := user.LoginUserRequest{Username: "test", Password: "abc"}
	_, status, err := uc.LoginUser(&request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, status)
}

func TestUseCaseLoginUserUnauthorized(t *testing.T) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("abc"), 14)
	uc := user.UseCase{
		Repository: &RepositoryMock{
			Errors:     nil,
			Conditions: map[string]interface{}{"password": string(hash)},
		},
	}
	request := user.LoginUserRequest{Username: "test", Password: "123"}
	_, status, err := uc.LoginUser(&request)
	assert.Error(t, err)
	assert.Equal(t, http.StatusUnauthorized, status)
}
