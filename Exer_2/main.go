package main

import (
	"example/Exer_1/golang-training/Exer_2/routes"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func main() {
	var router *chi.Mux
	router = routes.Routers(router)
	http.ListenAndServe(":8080", router)
}
