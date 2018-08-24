package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRoute()
	log.Fatal(http.ListenAndServe(":8111", r.mux))
}
