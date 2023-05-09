package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
