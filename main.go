package main

import (
	"example_day1/golang-training/database"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db := database.Init()
	defer db.Close()
	h := database.Handle{DB: db}
	r := chi.NewRouter()
	r.Get("/user/{id}", h.ShowUser)
	r.Post("/user/register", h.CreateUser)
	r.Put("/user/{id}", h.UpdateUser)

	log.Fatal(http.ListenAndServe(":8888", r))
}
