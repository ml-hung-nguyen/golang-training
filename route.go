package main

import (
	"baitapgo_ngay1/golang-training/middleware"
	"baitapgo_ngay1/golang-training/user"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

type Router struct {
	mux *chi.Mux
	db  *gorm.DB
}

func NewRouter() *Router {

	var route Router
	route.db = InitDB()
	route.mux = chi.NewRouter()

	uh := user.NewHandler(route.db)

	route.mux.Route("/user", func(r chi.Router) {
		r.Post("/login", uh.LoginUser)
		r.Post("/create", uh.RegisterUser)
		r.Get("/{id}", uh.GetUser)
		r.Put("/update", middleware.Authentication(uh.UpdateUser))
	})

	return &route
}
