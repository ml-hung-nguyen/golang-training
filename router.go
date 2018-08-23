package main

import (
	"golang-training/user"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type Router struct {
	Mux *chi.Mux
	DB  *gorm.DB
	// uh  *user.Handler
}

func NewRouter() *Router {
	var r Router
	r.DB = Connect()
	r.Mux = chi.NewRouter()
	// ur := user.NewRepository(r.DB)
	// uc := user.NewUseCase(ur)
	uh := user.NewHandler(r.DB)
	r.Mux.Route("/users", func(u chi.Router) {
		u.Get("/{id:[0-9]+}", uh.DetailUser)
		u.Post("/register", uh.RegisterUser)
		u.Post("/login", uh.Login)
		u.Put("/update/{id}", user.Authentication(uh.UpdateUser))
	})
	return &r
}
