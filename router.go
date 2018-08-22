package main

import (
	"golang-training/user"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type Router struct {
	Mux *chi.Mux
	DB  *gorm.DB
	uh  *user.Handler
}

func NewRouter() *Router {
	var r Router
	r.DB = Connect()
	r.Mux = chi.NewRouter()
	ur := user.NewRepository(r.DB)
	uc := user.NewUseCase(ur)
	r.uh = user.NewHandler(uc)

	r.Mux.Get("/users/{id}", r.uh.DetailUser)
	r.Mux.Post("/users/register", r.uh.RegisterUser)
	r.Mux.Post("/users/login", r.uh.Login)
	r.Mux.Put("/users/update/{id}", user.Authentication(r.uh.UpdateUser))
	return &r
}
