package main

import (
	file "baitapgo_ngay1/golang-training/file"
	"fmt"

	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func main() {
	db = file.InitDB()
	defer db.Close()
	h := file.Handler{
		DB:   db,
		User: &file.User{},
	}
	r := chi.NewRouter()
	fmt.Println("Starting Server")
	r.Get("/users/{id}", h.GetUser)
	r.Post("/users/adduser", h.CreateUser)
	r.Put("/users/{id}", h.UpdateUser)

	log.Fatal(http.ListenAndServe(":8089", r))
}
