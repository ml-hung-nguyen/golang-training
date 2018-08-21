package main

import (
	"golang-training/auth"
	"golang-training/db"
	"golang-training/handler"
	"golang-training/repository"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db.Connect()
	repo := repository.NewRepository(db.DB)
	h := handler.NewHandler(repo)
	r := chi.NewRouter()
	r.Get("/users/{id}", h.DetailUser)
	r.Post("/users/register", h.CreateUser)
	r.Put("/users/update/{id}", h.UpdateUser)
	r.Delete("/users/{id}", h.DeleteUser)
	r.Post("/login", h.LoginlUser)

	r.Get("/posts/{id}", auth.Authentication(h.DetailPost))
	r.Post("/posts/create", h.CreatePost)
	r.Put("/posts/update/{id}", h.UpdatePost)
	log.Fatal(http.ListenAndServe(":8888", r))
}
