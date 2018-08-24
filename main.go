package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRoute()
	// repo := NewRepository(DB)
	// handler := NewHandler(repo)
	// r := chi.NewRouter()
	// fmt.Println("Starting Server")
	// r.Post("/user/register", handler.CreateUser)
	// r.Get("/user/{id}", handler.GetUser)
	// r.Put("/user/{id}", handler.UpdateUser)
	log.Fatal(http.ListenAndServe(":8000", r.mux))
}
