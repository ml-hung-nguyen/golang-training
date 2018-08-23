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
}

func NewRouter() *Router {
	var r Router
	r.db = ConnectDB()
	r.mux = chi.NewRouter()

	uh := user.NewHandler(r.db)
	r.mux.Route("/user", func(cr chi.Router) {
		cr.Post("/login", uh.LoginUser)
		cr.Post("/register", uh.RegisterUser)
		cr.Get("/{id}", uh.GetUser)
		cr.Put("/update", middleware.Authentication(uh.UpdateUser))
	})

	return &r
}
