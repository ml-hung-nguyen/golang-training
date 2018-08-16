package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var router *chi.Mux

func routers() *chi.Mux {
	h := Handle{&Repository{db}}

	router.Get("/users/{id_user}", Authentication(h.DetailUser))
	router.Post("/users/register", h.CreateUser)
	router.Put("/users/update/{id_user}", Authentication(h.UpdateUser))
	router.Delete("/users/{id_user}", h.DeleteUser)
	router.Post("/login", h.LoginlUser)

	router.Get("/posts/{id_post}", h.DetailPost)
	router.Post("/posts/create", h.CreatePost)
	router.Put("/posts/update/{id_post}", h.UpdatePost)

	return router
}

func main() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	newConnect()
	routers()
	http.ListenAndServe(":8080", router)
}
