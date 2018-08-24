package routes

import (
	"example/Exer_1/golang-training/Exer_2/connect"
	"example/Exer_1/golang-training/Exer_2/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
)

func Routers(router *chi.Mux) *chi.Mux {
	var db *gorm.DB
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	db = connect.NewConnect(db)

	userHandler := user.NewUserHandler(db)

	router.Route("/user", func(r chi.Router) {
		r.Get("/{id_user:[0-9]+}", userHandler.DetailUserHandler)
		r.Put("/{id_user:[0-9]+}/update", user.Authentication(userHandler.UpdateUserHandler))
		r.Delete("/{id_user:[0-9]+}", userHandler.DeleteUserHandler)
	})

	router.Post("/register", userHandler.CreateUserHandler)
	router.Post("/login", userHandler.LoginlUser)

	return router
}
