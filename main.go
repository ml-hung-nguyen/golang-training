package main

import (
	"database/sql"
	"fmt"
	"golang-training/lib"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

var db *sql.DB

func main() {
	db := lib.NewConnect()
	defer db.Close()
	h := lib.Handler{DB: db}
	r := chi.NewRouter()
	fmt.Println("Starting Server")
	r.Get("/user/{id}", h.GetUser)
	r.Post("/user", h.CreateUser)
	r.Put("/user/{id}", h.UpdateUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}
