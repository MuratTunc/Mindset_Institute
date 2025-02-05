package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Constants for error and success messages
const (
	ErrInvalidRequestBody = "Invalid request body"
	ErrHashingPassword    = "Error hashing password"
	ErrInsertingUser      = "Error inserting user"
	ErrUserNotFound       = "User not found"
	ErrInvalidCredentials = "Invalid credentials"
	UserCreatedSuccess    = "User created successfully"
	LoginSuccess          = "Login successful"
)

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
