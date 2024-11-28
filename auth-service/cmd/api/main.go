package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const port = "4000"

func main() {
	err := godotenv.Load("./cmd/api/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	app := CreateApp()
	http.ListenAndServe(fmt.Sprintf(":%s", port), app)


}