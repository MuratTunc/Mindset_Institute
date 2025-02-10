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
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                 // Allow these HTTP methods for cross-origin requests
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, // Allowed headers in requests
		ExposedHeaders:   []string{"Link"},                                                    // Expose headers to the client
		AllowCredentials: true,                                                                // Allow cookies to be sent with cross-origin requests
		MaxAge:           300,                                                                 // Cache pre-flight requests for 5 minutes
	}))

	// Middleware
	mux.Use(middleware.Heartbeat("/ping")) // Basic health check endpoint (ping route)
	mux.Use(middleware.Recoverer)          // Automatically recover from panics and return a 500 status code
	mux.Use(middleware.Logger)             // Log all HTTP requests

	// POST routes
	mux.Post("/register", app.CreateCustomerHandler)        // Route to handle customer registration
	mux.Post("/login", app.LoginCustomerHandler)            // Route to handle customer login
	mux.Post("/update-password", app.UpdatePasswordHandler) // Route to update customer password (authentication required)

	// GET routes
	mux.Get("/health", app.HealthCheckHandler)                            // Custom health check endpoint to check if the service is up
	mux.Get("/get_all_customer", app.GetAllCustomerHandler)               // Route to get all customers
	mux.Get("/order-customers", app.OrderCustomersHandler)                // Route to order customers based on some criteria
	mux.Get("/activated-customers", app.GetActivatedCustomerNamesHandler) // Route to get activated customers
	mux.Get("/logged-in-customers", app.GetLoggedInCustomersHandler)      // Route to get logged-in customers
	mux.Get("/customer", app.GetCustomerHandler)                          // Route to get a specific customer by some criteria (e.g., ID)

	// PUT routes (used for updating data)
	mux.Put("/update-customer", app.UpdateCustomerHandler)         // Route to update customer information
	mux.Put("/deactivate-customer", app.DeactivateCustomerHandler) // Route to deactivate a customer account
	mux.Put("/activate-customer", app.ActivateCustomerHandler)     // Route to activate a customer account
	mux.Put("/update-email", app.UpdateEmailHandler)               // Route to update customer email
	mux.Put("/update-note", app.UpdateNoteHandler)                 // Route to update an existing customer note
	mux.Put("/insert-note", app.InsertNoteHandler)                 // Route to insert a new note into an existing customer note

	// DELETE routes
	mux.Delete("/delete-customer", app.DeleteCustomerHandler) // Route to delete a customer

	// Return the configured router to be used by the server
	return mux
}
