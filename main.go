package main

import (
	"log"
	"net/http"
	"os"
	"request-scoped-context-adv/handlers"
	"request-scoped-context-adv/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Middleware to add request ID and user data to context
	r.Use(middleware.RequestIDMiddleware)
	r.Use(middleware.UserMiddleware)

	// Handlers
	r.HandleFunc("/hello", handlers.HelloHandler).Methods("GET")
	r.HandleFunc("/process", handlers.ProcessHandler).Methods("POST")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
