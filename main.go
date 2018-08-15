package main

import (
	file "baitapgo_ngay1/golang-training/file"
	"database/sql"
	"fmt"

	"log"
	"net/http"

	"github.com/go-chi/chi"
)

var db *sql.DB

func main() {
	db = file.InitDB()
	defer db.Close()
	h := file.Handler{db, &file.User{}}
	r := chi.NewRouter()
	fmt.Println("Starting Server")
	r.Get("/users/{id}", h.GetUser)
	r.Post("/users/adduser", h.CreateUser)
	r.Put("/users/{id}", h.UpdateUser)

	log.Fatal(http.ListenAndServe(":8089", r))
}
