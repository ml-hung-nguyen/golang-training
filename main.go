package main

import (
	"golang-training/user"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	db := user.ConnectDB()
	defer db.Close()
	repo := user.NewRepository(db)
	h := user.NewHandler(repo)

	r := chi.NewRouter()
	r.Post("/user/login", h.LoginUser)
	r.Post("/user/register", h.RegisterUser)
	r.Get("/user/{id}", h.GetUser)
	r.Put("/user/update", user.Authentication(h.UpdateUser))
	r.Get("/posts", h.GetPostList)
	r.Get("/user/posts", user.Authentication(h.GetPosts))
	r.Post("/user/posts", user.Authentication(h.CreatePost))
	r.Put("/user/posts/{id}", user.Authentication(h.UpdatePost))
	r.Delete("/user/posts/{id}", user.Authentication(h.DeletePost))

	log.Fatal(http.ListenAndServe(":1709", r))
}
