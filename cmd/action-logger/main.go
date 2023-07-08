package main

import (
	"action-logger/pkg/handlers"
	"action-logger/pkg/middlewares"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/submit-logger", handlers.SubmitHandler)

	// Create a new LoggingMiddleware and wrap the existing handlers
	loggingMiddleware := middlewares.LoggingMiddleware(http.DefaultServeMux)

	// Use the LoggingMiddleware as the main handler
	mainHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggingMiddleware.ServeHTTP(w, r)
	})

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mainHandler))
}
