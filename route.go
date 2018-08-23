package main

import (
	"golang-training-example/user"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

//Route struct
type Route struct {
	mux *chi.Mux
	db  *gorm.DB
	uh  *user.Handler
}

//NewRoute func
func NewRoute() *Route {
	var r Route
	r.db = NewConnect()
	r.mux = chi.NewRouter()

	repo := user.NewRepository(r.db)
	uc := user.NewUseCase(repo)
	r.uh = user.NewHandler(uc)
	//r := chi.NewRouter()
	//fmt.Println("Starting Server")
	r.mux.Post("/user/register", r.uh.CreateUser)
	r.mux.Put("/user/{id}", r.uh.UpdateUser)
	r.mux.Get("/user/{id}", r.uh.GetUser)

	return &r
}
