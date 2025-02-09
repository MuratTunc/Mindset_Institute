package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
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
	mux.Get("/swagger/*", httpSwagger.WrapHandler)

	// Public Routes (No authentication required)
	mux.Post("/register", app.CreateUserHandler) // Handle user registration
	mux.Post("/login", app.LoginUserHandler)     // Handle user login

	// Protected Routes (Require JWT authentication)
	mux.Group(func(mux chi.Router) {
		mux.Use(AuthMiddleware) // Apply JWT authentication middleware

		mux.Get("/user", app.GetUserHandler)                    // Retrieve a user
		mux.Post("/update-password", app.UpdatePasswordHandler) // Password update (requires authentication)
		mux.Put("/update-user", app.UpdateUserHandler)          // Update user
		mux.Put("/deactivate-user", app.DeactivateUserHandler)  // Deactivate user by username
		mux.Put("/activate-user", app.ActivateUserHandler)      // Activate user
		mux.Put("/update-email", app.UpdateEmailHandler)        // Update user's email address
		mux.Put("/update-role", app.UpdateRoleHandler)          // Update user's role
		mux.Delete("/delete-user", app.DeleteUserHandler)       // Delete user

	})

	return mux
}
