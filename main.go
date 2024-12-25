package main

import (
	"log"
	"net/http"
	"weather-api/routes"
)

func main() {
	r := routes.SetupRoutes() // Setup all routes
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r)) // Start server on port 8080
}
