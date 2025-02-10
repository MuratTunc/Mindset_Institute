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

	mux.Post("/register", app.CreateCustomerHandler)        // Handle customer registration
	mux.Post("/login", app.LoginCustomerHandler)            // Handle customer login
	mux.Post("/update-password", app.UpdatePasswordHandler) // Password update (requires authentication)

	mux.Get("/health", app.HealthCheckHandler) // Custom health check endpoint
	mux.Get("/get_all_customer", app.GetAllCustomerHandler)
	mux.Get("/order-customers", app.OrderCustomersHandler)
	mux.Get("/activated-customers", app.GetActivatedCustomerNamesHandler)
	mux.Get("/logged-in-customers", app.GetLoggedInCustomersHandler)
	mux.Get("/customer", app.GetCustomerHandler) // Retrieve a customer by ID

	mux.Put("/update-customer", app.UpdateCustomerHandler)         // Update customer by ID
	mux.Put("/deactivate-customer", app.DeactivateCustomerHandler) // Deactivate customer by ID
	mux.Put("/activate-customer", app.ActivateCustomerHandler)     // Activate customer by ID
	mux.Put("/update-email", app.UpdateEmailHandler)               // Update customer's email address
	mux.Put("/update-note", app.UpdateNoteHandler)
	mux.Put("/insert-note", app.InsertNoteHandler) // Route to insert new note into an existing one

	mux.Delete("/delete-customer", app.DeleteCustomerHandler) // Delete customer

	return mux
}
