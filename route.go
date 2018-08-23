package main

import (
	"golang-training/midleware"
	"golang-training/user"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

//Route struct
type Route struct {
	mux *chi.Mux
	db  *gorm.DB
	//h   *user.Handler
}

//NewRoute func
func NewRoute() *Route {
	var r Route
	r.db = Init()
	r.mux = chi.NewRouter()

	repo := user.NewRepository(r.db)
	uc := user.NewUseCase(repo)
	h := user.NewHandler(uc)

	r.mux.Post("/user/login", h.LoginUser)
	r.mux.Post("/user/register", h.CreateUser)
	r.mux.Get("/user/{id}", h.ShowUser)
	r.mux.Put("/user/update", midleware.Authentication(h.UpdateUser))

	return &r
}
