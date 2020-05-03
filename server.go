package main

import (
	"net/http"

	"github.com/anujkmaurya/golang-fifa-world-cup-web-service/handlers"

	"github.com/anujkmaurya/golang-fifa-world-cup-web-service/data"
)

func main() {
	data.PrintUsage()

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/winners", handlers.WinnersHandler)
	http.ListenAndServe(":8000", nil)
}
