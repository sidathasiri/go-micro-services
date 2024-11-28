package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "3000"


func main() {
	app := CreateApp()
	http.ListenAndServe(fmt.Sprintf(":%s", port), app)
	log.Println("Server started on port ", port)
}