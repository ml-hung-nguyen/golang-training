package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var router *chi.Mux

func routers() *chi.Mux {
	h := Handler{db, &User{}}

	router.Get("/users/{id_user}", h.Detail)
	router.Post("/users/register", h.Create)
	router.Put("/users/update/{id_user}", h.Update)

	return router
}

func main() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	Connect()
	routers()
	http.ListenAndServe(":8080", router)
}
