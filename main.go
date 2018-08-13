package main

import (
	"golang-training/user"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db := ConnectDB()
	defer db.Close()

	user := user.UserHandler{DB: db}

	r := chi.NewRouter()
	r.Post("/user/register", user.RegisterHandler)
	r.Get("/user/{id}", user.GetHandler)
	r.Put("/user/{id}", user.UpdateHandler)
	log.Fatal(http.ListenAndServe(":1709", r))
}
