package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Constants for error and success messages
const (
	ErrInvalidRequestBody  = "Invalid request body"
	ErrHashingPassword     = "Error hashing password"
	ErrInsertingCustomer   = "Error inserting customer"
	ErrCustomerNotFound    = "Customer not found"
	ErrInvalidCredentials  = "Invalid credentials"
	CustomerCreatedSuccess = "Customer created successfully"
	CustomerUpdatedSuccess = "Customer updated successfully"
	CustomerDeletedSuccess = "Customer deleted successfully"
	LoginSuccess           = "Login successful"
)

// HealthCheckHandler checks the database connection using GORM
func (app *Config) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HealthCheckHandler is Called")
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

// CreateCustomerHandler inserts a new customer with a hashed password using GORM
func (app *Config) CreateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure MailAddress is provided
	if customer.MailAddress == "" {
		http.Error(w, "Mail address cannot be empty", http.StatusBadRequest)
		return
	}

	// Check if customer already exists (by customername OR mail address)
	var existingCustomer Customer
	if err := app.DB.Where("customername = ? OR mail_address = ?", customer.Customername, customer.MailAddress).First(&existingCustomer).Error; err == nil {
		http.Error(w, "Customer already exists", http.StatusConflict)
		return
	}

	// Hash the password before saving
	hashedPassword, err := app.HashPassword(customer.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	customer.Password = hashedPassword

	// Set Activated = true for every new customer
	customer.Activated = true

	// Insert customer into database using GORM
	result := app.DB.Create(&customer)
	if result.Error != nil {
		http.Error(w, "Error inserting customer", http.StatusInternalServerError)
		return
	}

	// Send response (without JWT token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":     "Customer created successfully",
		"mailAddress": customer.MailAddress, // Include mail address in the response
	})
}

// LoginCustomerHandler verifies customer credentials using GORM
func (app *Config) LoginCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	var storedCustomer Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find customer in DB
	result := app.DB.Where("customername = ?", customer.Customername).First(&storedCustomer)
	if result.Error != nil {
		http.Error(w, "Customer not found", http.StatusUnauthorized)
		return
	}

	// Compare passwords
	if !app.CheckPassword(storedCustomer.Password, customer.Password) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Update login_status to true
	storedCustomer.LoginStatus = true
	app.DB.Save(&storedCustomer)

	// Send response (without JWT token)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":     "Login successful",
		"loginStatus": "true",
	})
}

// UpdatePasswordHandler updates a customer's password if they exist
func (app *Config) UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Customername string `json:"customername"`
		NewPassword  string `json:"new_password"`
	}

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, ErrInvalidRequestBody, http.StatusBadRequest)
		return
	}

	// Log the incoming request for debugging
	fmt.Println("Received request to update password for customername:", requestData.Customername)

	// Find the customer by customername
	var customer Customer
	result := app.DB.Where("customername = ?", requestData.Customername).First(&customer)

	// Log the result of the query to help debug
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, ErrCustomerNotFound, http.StatusNotFound)
			fmt.Println("Customer not found:", requestData.Customername) // Add log for failed query
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Log the found customer for debugging
	fmt.Println("Found customer:", customer)

	// Hash new password
	hashedPassword, err := app.HashPassword(requestData.NewPassword)
	if err != nil {
		http.Error(w, ErrHashingPassword, http.StatusInternalServerError)
		return
	}

	// Update the password
	customer.Password = hashedPassword
	app.DB.Save(&customer)

	// Log successful password update
	fmt.Println("Password updated for customer:", requestData.Customername)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Password updated successfully")
}

// GetCustomerHandler retrieves a customer by ID
func (app *Config) GetCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	id := r.URL.Query().Get("id")

	// Fetch the customer by ID from the database
	result := app.DB.First(&customer, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, ErrCustomerNotFound, http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if the MailAddress is empty and log a warning or handle appropriately
	if customer.MailAddress == "" {
		// You can log this if necessary or handle the empty field case
		fmt.Println("Warning: Customer has no MailAddress.")
	}

	// Respond with customer data in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

// UpdateCustomerHandler updates customer details
func (app *Config) UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Get the customer ID from the URL path parameter
	id := chi.URLParam(r, "id")

	var customer Customer
	result := app.DB.First(&customer, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, ErrCustomerNotFound, http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, ErrInvalidRequestBody, http.StatusBadRequest)
		return
	}

	// If password is being updated, hash it
	if customer.Password != "" {
		hashedPassword, err := app.HashPassword(customer.Password)
		if err != nil {
			http.Error(w, ErrHashingPassword, http.StatusInternalServerError)
			return
		}
		customer.Password = hashedPassword
	}

	app.DB.Save(&customer)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, CustomerUpdatedSuccess)
}

func (app *Config) DeactivateCustomerHandler(w http.ResponseWriter, r *http.Request) {

	// Extract customer ID by splitting the path
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 2 {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	// The customer ID is expected to be the second segment of the URL path
	id := segments[len(segments)-1]
	fmt.Println("Extracted Customer ID:", id)

	// Check if the ID is valid
	if id == "" {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	//  Finds the customer in the database.
	var customer Customer
	result := app.DB.First(&customer, id)
	if result.Error != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	// Set Activated to false
	customer.Activated = false

	// Update the customer in the database
	if err := app.DB.Save(&customer).Error; err != nil {
		http.Error(w, "Failed to deactivate customer", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":     "Customer deactivated successfully",
		"customer_id": id,
	})
}

func (app *Config) ActivateCustomerHandler(w http.ResponseWriter, r *http.Request) {

	// Extract customer ID by splitting the path
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 2 {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	// The customer ID is expected to be the second segment of the URL path
	id := segments[len(segments)-1]
	fmt.Println("Extracted Customer ID:", id)

	// Check if the ID is valid
	if id == "" {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	//  Finds the customer in the database.
	var customer Customer
	result := app.DB.First(&customer, id)
	if result.Error != nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	// Set Activated to false
	customer.Activated = true

	// Update the customer in the database
	if err := app.DB.Save(&customer).Error; err != nil {
		http.Error(w, "Failed to activate customer", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":     "Customer activated successfully",
		"customer_id": id,
	})
}

// UpdateEmailHandler updates the customer's email address by ID
func (app *Config) UpdateEmailHandler(w http.ResponseWriter, r *http.Request) {
	// Extract customer ID by splitting the path
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 2 {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	// The customer ID is expected to be the second segment of the URL path
	id := segments[len(segments)-1]
	fmt.Println("Extracted Customer ID:", id)

	// Check if the ID is valid
	if id == "" {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	// Parse the new email address from the request body (assuming JSON format)
	var requestData struct {
		NewEmail string `json:"new_email"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure the new email is not empty
	if requestData.NewEmail == "" {
		http.Error(w, "New email is required", http.StatusBadRequest)
		return
	}

	// Check if the new email is valid (basic format validation)
	if !isValidEmail(requestData.NewEmail) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Find the customer by ID
	var customer Customer
	result := app.DB.Where("id = ?", id).First(&customer)

	// If the customer is not found
	if result.RowsAffected == 0 {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	// Update the customer's email address
	customer.MailAddress = requestData.NewEmail
	if err := app.DB.Save(&customer).Error; err != nil {
		http.Error(w, "Failed to update email", http.StatusInternalServerError)
		return
	}

	// Log the email update action (optional)
	fmt.Printf("Customer %s (ID: %s) updated their email\n", customer.Customername, id)

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Email updated successfully")
}

// Helper function to validate email format (simple validation)
func isValidEmail(email string) bool {
	// You can use a more robust regex or library for email validation
	// Here, a basic validation for a common email format is used
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// UpdateNoteHandler updates the Note field for a customer
func (app *Config) UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Customername string `json:"customername"`
		Note         string `json:"note"`
	}

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find customer in DB
	var customer Customer
	result := app.DB.Where("customername = ?", request.Customername).First(&customer)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Customer not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Update the Note field
	customer.Note = request.Note
	if err := app.DB.Save(&customer).Error; err != nil {
		http.Error(w, "Failed to update note", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Note updated successfully",
	})
}

// DeleteCustomerHandler deletes a customer by ID
func (app *Config) DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Extract customer ID by splitting the path
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 2 {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	// The customer ID is expected to be the second segment of the URL path
	id := segments[len(segments)-1]
	fmt.Println("Extracted Customer ID:", id)

	// Ensure the customer ID is valid
	if id == "" {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	// Find the customer to delete by ID
	var customer Customer
	result := app.DB.Where("id = ?", id).Delete(&customer)

	// If no rows were affected, the customer was not found
	if result.RowsAffected == 0 {
		http.Error(w, ErrCustomerNotFound, http.StatusNotFound)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Customer deleted successfully")
}
