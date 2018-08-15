package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var route *chi.Mux

func router() *chi.Mux {
	h := HandlerUser{db, &User{}}
	route.Get("/users/{id}", h.Detail)
	route.Post("/users/register", h.Create)
	route.Put("/users/update/{id}", h.Update)
	return route
}

func main() {
	Connect()
	route = chi.NewRouter()
	route.Use(middleware.Recoverer)
	router()
	log.Println("Starting server:")
	http.ListenAndServe(":8888", route)
}
