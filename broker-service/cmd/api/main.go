package main

import (
	"fmt"
	"net/http"
)

const port = "3000"


func main() {
	app := CreateApp()
	http.ListenAndServe(fmt.Sprintf(":%s", port), app)
}