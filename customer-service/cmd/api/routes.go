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
			"http://localhost:3000", // Allow requests from localhost our React frontend web-app
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Middleware
	mux.Use(middleware.Heartbeat("/ping")) // Basic health check
	mux.Use(middleware.Recoverer)          // Recover from panics gracefully
	mux.Use(middleware.Logger)             // Log all requests

	// Custom health check endpoint
	mux.Get("/health", app.HealthCheckHandler)

	// Public Routes (No authentication required)
	mux.Post("/register", app.CreateCustomerHandler) // Handle customer registration
	mux.Post("/login", app.LoginCustomerHandler)     // Handle customer login
	mux.Get("/get_all_customer", app.GetAllCustomerHandler)

	mux.Get("/customer", app.GetCustomerHandler)                        // Retrieve a customer by ID
	mux.Post("/update-password", app.UpdatePasswordHandler)             // Password update (requires authentication)
	mux.Put("/update-customer/{id}", app.UpdateCustomerHandler)         // Update customer by ID
	mux.Put("/deactivate-customer/{id}", app.DeactivateCustomerHandler) // Deactivate customer by ID
	mux.Put("/activate-customer/{id}", app.ActivateCustomerHandler)     // Activate customer by ID
	mux.Put("/update-email/{id}", app.UpdateEmailHandler)               // Update customer's email address
	mux.Put("/update-note", app.UpdateNoteHandler)
	mux.Put("/insert-note", app.InsertNoteHandler)                 // Route to insert new note into an existing one
	mux.Delete("/delete-customer/{id}", app.DeleteCustomerHandler) // Delete customer by ID

	return mux
}
