package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()
	log.Fatal(http.ListenAndServe(":1709", r.mux))
}
