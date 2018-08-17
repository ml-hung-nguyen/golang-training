package main

import (
	"example_day1/golang-training/file"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db := file.Init()
	defer db.Close()
	repo := file.NewRepository(db)
	h := file.NewHandler(repo)
	r := chi.NewRouter()
	r.Get("/user/{id}", h.ShowUser)
	r.Post("/user/register", h.CreateUser)
	r.Put("/user/{id}", h.UpdateUser)

	log.Fatal(http.ListenAndServe(":8112", r))
}
