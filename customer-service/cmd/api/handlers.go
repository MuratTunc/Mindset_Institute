package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
func (app *Config) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, ErrInvalidRequestBody, http.StatusBadRequest)
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

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, UserCreatedSuccess)
}

// LoginUserHandler verifies user credentials using GORM
func (app *Config) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	var storedUser User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, ErrInvalidRequestBody, http.StatusBadRequest)
		return
	}

	// Find user in DB
	result := app.DB.Where("username = ?", user.Username).First(&storedUser)
	if result.Error != nil {
		http.Error(w, ErrUserNotFound, http.StatusUnauthorized)
		return
	}

	// Compare passwords
	if !app.CheckPassword(storedUser.Password, user.Password) {
		http.Error(w, ErrInvalidCredentials, http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, LoginSuccess)
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

	result := app.DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, ErrUserNotFound, http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUserHandler updates user details
func (app *Config) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	id := r.URL.Query().Get("id")

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
	var user User
	id := r.URL.Query().Get("id")

	result := app.DB.Delete(&user, id)
	if result.RowsAffected == 0 {
		http.Error(w, ErrUserNotFound, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, UserDeletedSuccess)
}
