package user_test

import (
	"bytes"
	"golang-training/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	h := &user.Handler{
		UseCase: &UseCaseMock{},
	}
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer([]byte(`{"username": "QQ","password": "abc","fullname": "QO"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Post("/user/register", h.RegisterUser)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
}

func TestRegisterUserNoBody(t *testing.T) {
	h := &user.Handler{
		UseCase: &UseCaseMock{},
	}
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer([]byte(``)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Post("/user/register", h.RegisterUser)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
}

func TestRegisterUserInvalidRequest(t *testing.T) {
	h := &user.Handler{
		UseCase: &UseCaseMock{},
	}
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer([]byte(`"username": "QQ","password": "abc","fullname": "QO"`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Post("/user/register", h.RegisterUser)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}

func TestRegisterUserFailValidate(t *testing.T) {
	h := &user.Handler{
		UseCase: &UseCaseMock{},
	}
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer([]byte(`{"password": "abc","fullname": "QO"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Post("/user/register", h.RegisterUser)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Result().StatusCode)
}

func TestRegisterUserFailSQL(t *testing.T) {
	h := &user.Handler{
		UseCase: &UseCaseMock{
			Errors: gorm.ErrInvalidSQL,
			Status: http.StatusInternalServerError,
		},
	}
	req, err := http.NewRequest("POST", "/user/register", bytes.NewBuffer([]byte(`{"username": "QQ","password": "abc","fullname": "QO"}`)))
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}
	w := httptest.NewRecorder()
	r := chi.NewRouter()
	r.Post("/user/register", h.RegisterUser)

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}
