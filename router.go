package main

import (
	"golang-training/user"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type Route struct {
	mux *chi.Mux
	db  *gorm.DB
	uh  *user.Handler
}

//NewRoute func
func NewRoute() *Route {
	var r Route
	r.db = Connect()
	r.mux = chi.NewRouter()

	ur := user.NewRepository(r.db)
	uc := user.NewUseCase(ur)
	r.uh = user.NewHandler(uc)

	r.mux.Post("/users/login", r.uh.LoginUser)
	r.mux.Post("/users/register", r.uh.RegisterUser)
	r.mux.Get("/users/{id}", r.uh.GetUser)
	r.mux.Put("/users/update", user.Authentication(r.uh.UpdateUser))

	return &r
}
