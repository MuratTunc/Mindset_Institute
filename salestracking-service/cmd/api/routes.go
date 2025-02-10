package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Define routes for the application
func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// CORS Middleware
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:3000", // Allow requests from localhost, where our React frontend web app is running
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                 // Allowed HTTP methods for cross-origin requests
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // Allowed request headers
		ExposedHeaders:   []string{"Link"},                                                    // Expose headers to the client
		AllowCredentials: true,                                                                // Allow credentials (cookies) to be sent with cross-origin requests
		MaxAge:           300,                                                                 // Cache pre-flight requests for 5 minutes
	}))

	// Middleware
	mux.Use(middleware.Heartbeat("/ping")) // Health check route that returns 200 OK if the server is up
	mux.Use(middleware.Recoverer)          // Recovers from panics and avoids server crashes
	mux.Use(middleware.Logger)             // Logs all incoming HTTP requests

	// Custom health check endpoint
	mux.Get("/health", app.HealthCheckHandler) // Route to check if the service is healthy and running

	// POST routes for creating resources
	mux.Post("/insert-sale", app.InsertSaleHandler) // Route to insert a new sale record into the system

	// DELETE routes for removing resources
	mux.Delete("/delete-sale", app.DeleteSaleHandler) // Route to delete an existing sale record from the system

	// PUT routes for updating existing resources
	mux.Put("/update-incommunication", app.UpdateInCommunicationHandler) // Route to update the "in communication" status of a sale
	mux.Put("/update-deal", app.UpdateDealHandler)                       // Route to update the "deal" status of a sale
	mux.Put("/update-closed", app.UpdateClosedHandler)                   // Route to update the "closed" status of a sale

	// Return the configured router to be used by the server
	return mux
}
