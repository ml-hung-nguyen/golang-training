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

	u := user.UserHandler{
		DB:   db,
		User: &user.User{},
	}

	r := chi.NewRouter()
	r.Post("/user/login", u.LoginHandler)
	r.Post("/user/register", u.RegisterHandler)
	r.Get("/user/{id}", u.GetHandler)
	r.Put("/user/update", user.Authentication(u.UpdateHandler))
	log.Fatal(http.ListenAndServe(":1709", r))
}
