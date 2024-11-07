package main

import (
	"log"
	"net/http"

	"golangAssignment/handler"
	validator "golangAssignment/helper"
	"golangAssignment/middleware"
	"golangAssignment/service"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router and the service
	r := mux.NewRouter()

	// Initialize validator
	validate := validator.NewValidator()

	// Initialize the service with the validator
	svc := service.NewService(validate)

	// Initialize the handler
	h := handler.NewHandler(svc)

	// Set up routes
	r.HandleFunc("/validate", h.HandleRequest).Methods("POST")

	// Apply middleware
	r.Use(middleware.LatencyLogger)

	// Start the server
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
