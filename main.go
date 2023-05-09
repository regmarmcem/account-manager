package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/regmarmcem/account-manager/api"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file not found")
	}

	dbHost := os.Getenv("DB_HOSTNAME")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	conn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUser, dbPass)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("failed to connect database")
	}

	r := api.NewRouter(db)
	fmt.Println("web server starting...")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
