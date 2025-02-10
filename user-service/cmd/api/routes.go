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
	mux.Get("/health", app.HealthCheckHandler)     // Route to check if the service is healthy and running
	mux.Get("/swagger/*", httpSwagger.WrapHandler) // Route to serve Swagger documentation (API documentation)

	// Public Routes (No authentication required)
	mux.Post("/register", app.CreateUserHandler) // Route to register a new user (no authentication needed)
	mux.Post("/login", app.LoginUserHandler)     // Route to log in an existing user (no authentication needed)

	// Protected Routes (Require JWT authentication)
	mux.Group(func(mux chi.Router) {
		mux.Use(AuthMiddleware) // Apply JWT authentication middleware to all routes within this group

		// User management routes (authentication required)
		mux.Get("/user", app.GetUserHandler)                    // Route to retrieve a user by their ID
		mux.Post("/update-password", app.UpdatePasswordHandler) // Route to update the user's password (authentication required)
		mux.Put("/update-user", app.UpdateUserHandler)          // Route to update user information (authentication required)
		mux.Put("/deactivate-user", app.DeactivateUserHandler)  // Route to deactivate a user by username (authentication required)
		mux.Put("/activate-user", app.ActivateUserHandler)      // Route to activate a deactivated user (authentication required)
		mux.Put("/update-email", app.UpdateEmailHandler)        // Route to update the user's email address (authentication required)
		mux.Put("/update-role", app.UpdateRoleHandler)          // Route to update the user's role (authentication required)
		mux.Delete("/delete-user", app.DeleteUserHandler)       // Route to delete a user (authentication required)
	})

	// Return the configured router to be used by the server
	return mux
}
