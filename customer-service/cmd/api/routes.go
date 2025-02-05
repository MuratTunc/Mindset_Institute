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
	mux.Use(middleware.Heartbeat("/ping")) // Health check endpoint
	mux.Use(middleware.Recoverer)          // Recover from panics gracefully
	mux.Use(middleware.Logger)             // Log all requests

	// Routes for authentication
	mux.Post("/register", app.CreateUserHandler)            // Handle registration
	mux.Post("/login", app.LoginUserHandler)                // Handle login
	mux.Post("/update-password", app.UpdatePasswordHandler) // New password update route

	// CRUD operations for users
	mux.Get("/user", app.GetUserHandler)       // Retrieve a user by ID (query parameter)
	mux.Put("/user", app.UpdateUserHandler)    // Update user by ID (query parameter)
	mux.Delete("/user", app.DeleteUserHandler) // Delete user by ID (query parameter)

	return mux
}
