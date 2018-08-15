package main

import (
	"github.com/at-hungnguyen2/golang-training/db"
	"github.com/at-hungnguyen2/golang-training/handler"
	"github.com/at-hungnguyen2/golang-training/repository"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	db.NewConnect()
	repo := repository.NewRepository(db.DB)
	handler := handler.NewHandler(repo)
	r := chi.NewRouter()
	r.Post("/user/register", handler.CreateUser)
	log.Fatal(http.ListenAndServe(":1709", r))
}
