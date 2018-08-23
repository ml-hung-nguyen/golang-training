package main

import (
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	route := NewRouter()
	http.ListenAndServe(":8089", route.mux)
}
