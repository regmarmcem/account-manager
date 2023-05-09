package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file not found")
	}

	r := chi.NewRouter()
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
