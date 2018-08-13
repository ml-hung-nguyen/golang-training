package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (u *UserHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		w.WriteHeader(400)
		return
	}
	if len(body) < 1 {
		json.NewEncoder(w).Encode(ErrorResponse{Message: "No body"})
		w.WriteHeader(400)
		return
	}

	err = json.Unmarshal(body, &user)
	fmt.Println(user)
	if u.DB.NewRecord(user) {
		u.DB.Create(&user)
		fmt.Println("Created")
	}
	json.NewEncoder(w).Encode(user)
	return
}

func (u *UserHandler) GetHandler(w http.ResponseWriter, r *http.Request) {

}

func (u *UserHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {

}
