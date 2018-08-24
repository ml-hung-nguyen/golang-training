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
	h := user.NewHandler(r.db)
	r.mux.Route("/user", func(c chi.Router) {
		c.Post("/login", h.LoginUser)
		c.Post("/register", h.CreateUser)
		c.Get("/{id:[0-9]+}", h.ShowUser)
		c.Put("/update", midleware.Authentication(h.UpdateUser))
	})
	// r.mux.Post("/user/login", h.LoginUser)
	// r.mux.Post("/user/register", h.CreateUser)
	// r.mux.Get("/user/{id}", h.ShowUser)
	// r.mux.Put("/user/update", midleware.Authentication(h.UpdateUser))
	return &r
}
