package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Constants for error and success messages
const (
	ErrInvalidRequestBody = "Invalid request body"
	ErrHashingPassword    = "Error hashing password"
	ErrInsertingUser      = "Error inserting user"
	ErrUserNotFound       = "User not found"
	ErrInvalidCredentials = "Invalid credentials"
	UserCreatedSuccess    = "User created successfully"
	UserUpdatedSuccess    = "User updated successfully"
	UserDeletedSuccess    = "User deleted successfully"
	LoginSuccess          = "Login successful"
)

// Secret key for JWT signing
var jwtSecret = []byte(JWTSecret)

// GenerateJWT creates a JWT token for a user
func GenerateJWT(username, role string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// AuthMiddleware verifies JWT tokens
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Extract token from "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader { // No "Bearer " prefix found
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		// Parse and verify JWT
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add username and role to request context (optional)
		r.Header.Set("X-Username", claims["username"].(string))
		r.Header.Set("X-Role", claims["role"].(string))

		next.ServeHTTP(w, r) // Call the next handler
	})
}

// HealthCheckHandler checks the database connection using GORM
func (app *Config) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	sqlDB, err := app.DB.DB() // Get *sql.DB from *gorm.DB
	if err != nil {
		http.Error(w, "Failed to get database instance", http.StatusInternalServerError)
		return
	}

	// Execute a lightweight query to check DB connectivity
	err = sqlDB.Ping()
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// HashPassword hashes a password using bcrypt
func (app *Config) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a hashed password with a plain one
func (app *Config) CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// CreateUserHandler inserts a new user with a hashed password using GORM
// CreateUserHandler inserts a new user with a hashed password using GORM
func (app *Config) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, ErrInvalidRequestBody, http.StatusBadRequest)
		return
	}

	// Ensure MailAddress is provided
	if user.MailAddress == "" {
		http.Error(w, "Mail address cannot be empty", http.StatusBadRequest)
		return
	}

	// Check if user already exists (by username OR mail address)
	var existingUser User
	if err := app.DB.Where("username = ? OR mail_address = ?", user.Username, user.MailAddress).First(&existingUser).Error; err == nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	// Hash the password before saving
	hashedPassword, err := app.HashPassword(user.Password)
	if err != nil {
		http.Error(w, ErrHashingPassword, http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	// Insert user into database using GORM
	result := app.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, ErrInsertingUser, http.StatusInternalServerError)
		return
	}

	// Generate JWT token
	token, err := GenerateJWT(user.Username, user.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":     UserCreatedSuccess,
		"token":       token,
		"mailAddress": user.MailAddress, // Include mail address in the response
	})
}

// LoginUserHandler verifies user credentials using GORM
func (app *Config) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	var storedUser User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find user in DB
	result := app.DB.Where("username = ?", user.Username).First(&storedUser)
	if result.Error != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Compare passwords
	if !app.CheckPassword(storedUser.Password, user.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := GenerateJWT(storedUser.Username, storedUser.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Send token to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// UpdatePasswordHandler updates a user's password if they exist
func (app *Config) UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Username    string `json:"username"`
		NewPassword string `json:"new_password"`
	}

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, ErrInvalidRequestBody, http.StatusBadRequest)
		return
	}

	// Find the user
	var user User
	result := app.DB.Where("username = ?", requestData.Username).First(&user)
	if result.Error != nil {
		http.Error(w, ErrUserNotFound, http.StatusNotFound)
		return
	}

	// Hash new password
	hashedPassword, err := app.HashPassword(requestData.NewPassword)
	if err != nil {
		http.Error(w, ErrHashingPassword, http.StatusInternalServerError)
		return
	}

	// Update the password
	user.Password = hashedPassword
	app.DB.Save(&user)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Password updated successfully")
}

// GetUserHandler retrieves a user by ID
func (app *Config) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	id := r.URL.Query().Get("id")

	// Fetch the user by ID from the database
	result := app.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, ErrUserNotFound, http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if the MailAddress is empty and log a warning or handle appropriately
	if user.MailAddress == "" {
		// You can log this if necessary or handle the empty field case
		fmt.Println("Warning: User has no MailAddress.")
	}

	// Respond with user data in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// UpdateUserHandler updates user details
func (app *Config) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL path parameter
	id := chi.URLParam(r, "id")

	var user User
	result := app.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, ErrUserNotFound, http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, ErrInvalidRequestBody, http.StatusBadRequest)
		return
	}

	// If password is being updated, hash it
	if user.Password != "" {
		hashedPassword, err := app.HashPassword(user.Password)
		if err != nil {
			http.Error(w, ErrHashingPassword, http.StatusInternalServerError)
			return
		}
		user.Password = hashedPassword
	}

	app.DB.Save(&user)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, UserUpdatedSuccess)
}

// DeleteUserHandler deletes a user by ID
func (app *Config) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from the URL
	id := chi.URLParam(r, "id")

	// Ensure that the ID is valid (non-empty)
	if id == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Find the user to delete by ID
	var user User
	result := app.DB.Where("id = ?", id).Delete(&user)

	// If no rows were affected, the user was not found
	if result.RowsAffected == 0 {
		http.Error(w, ErrUserNotFound, http.StatusNotFound)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User deleted successfully")
}
