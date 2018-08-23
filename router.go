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

	r.uh = user.NewHandler(r.db)
	r.mux.Route("/user", func(cr chi.Router) {
		cr.Post("/login", r.uh.LoginUser)
		cr.Post("/register", r.uh.RegisterUser)
		cr.Get("/{id}", r.uh.GetUser)
		cr.Put("/update", middleware.Authentication(r.uh.UpdateUser))
	})

	return &r
}
