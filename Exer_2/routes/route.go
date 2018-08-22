package routes

import (
	"example/Exer_2/connect"
	"example/Exer_2/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
)

func Routers(router *chi.Mux) *chi.Mux {
	var db *gorm.DB
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	db = connect.NewConnect(db)

	conRepo := connect.NewConRepo(db)
	conUC := connect.NewConUC(conRepo)
	conHandel := connect.NewConHandle(conUC)

	userRepo := user.UserRepository{DB: db}
	userUC := user.UserUseCase{&userRepo}
	userHandel := user.UserHandle{&userUC}

	router.Get("/testdatabase", conHandel.ConnectHandle)
	router.Get("/users/{id_user}", userHandel.DetailUserHandle)
	router.Post("/users/register", userHandel.CreateUserHandle)
	router.Put("/users/{id_user}/update", user.Authentication(userHandel.UpdateUserHandle))
	router.Delete("/users/{id_user}", userHandel.DeleteUserHandle)
	router.Post("/login", userHandel.LoginlUser)

	return router
}
