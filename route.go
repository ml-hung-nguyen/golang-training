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
	uh  *user.Handler
}

func NewRouter() *Router {

	var route Router
	route.db = InitDB()
	route.mux = chi.NewRouter()

	ur := user.NewRepository(route.db)
	uc := user.NewUseCase(ur)
	route.uh = user.NewHandler(uc)

	route.mux.Post("/user/login", route.uh.LoginUser)
	route.mux.Post("/user/create", route.uh.RegisterUser)
	route.mux.Get("/user/{id}", route.uh.GetUser)
	route.mux.Put("/user/update", middleware.Authentication(route.uh.UpdateUser))

	return &route
}
