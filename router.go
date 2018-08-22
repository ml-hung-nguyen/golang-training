package main

import (
	"golang-training/middleware"
	"golang-training/user"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type Router struct {
	mux *chi.Mux
	db  *gorm.DB
	uh  *user.Handler
}

func NewRouter() *Router {
	var r Router
	r.db = ConnectDB()
	r.mux = chi.NewRouter()

	ur := user.NewRepository(r.db)
	uc := user.NewUseCase(ur)
	r.uh = user.NewHandler(uc)

	r.mux.Post("/user/login", r.uh.LoginUser)
	r.mux.Post("/user/register", r.uh.RegisterUser)
	r.mux.Get("/user/{id}", r.uh.GetUser)
	r.mux.Put("/user/update", middleware.Authentication(r.uh.UpdateUser))

	return &r
}
