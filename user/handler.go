package user

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (u *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandler) GetHandler(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {

}
