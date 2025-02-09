package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

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

// GetLoggedInCustomersHandler returns all customers with LoginStatus set to true
// @Summary Get all customers with active login status
// @Description This endpoint retrieves all customers whose LoginStatus is set to true
// @Tags customers
// @Accept  json
// @Produce  json
// @Success 200 {array} Customer "List of active customers"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 404 {object} ErrorResponse "No customers found with active login status"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /customers/logged-in [get]
func (app *Config) GetLoggedInCustomersHandler(w http.ResponseWriter, r *http.Request) {
	var customers []Customer

	// Fetch customers with LoginStatus = true
	result := app.DB.Where("login_status = ?", true).Find(&customers)
	if result.Error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if any customers were found
	if len(customers) == 0 {
		http.Error(w, "No customers with active login status", http.StatusNotFound)
		return
	}

	// Return the list of customers
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// GetActivatedCustomerNamesHandler returns the names of customers who are activated
// @Summary Get names of activated customers
// @Description This endpoint retrieves the names of customers who are activated
// @Tags customers
// @Accept  json
// @Produce  json
// @Success 200 {array} string "List of activated customer names"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 404 {object} ErrorResponse "No activated customers found"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /customers/activated/names [get]
func (app *Config) GetActivatedCustomerNamesHandler(w http.ResponseWriter, r *http.Request) {
	// Declare a slice to hold customer names
	var customerNames []string

	// Fetch only customers where 'Activated' is true
	// Use Model(&Customer) to specify the table
	result := app.DB.Model(&Customer{}).Where("activated = ?", true).Pluck("customername", &customerNames)
	if result.Error != nil {
		// Log the error to see what's going wrong
		log.Printf("Error fetching activated customers: %v", result.Error)
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Check if we have no activated customers
	if len(customerNames) == 0 {
		http.Error(w, "No activated customers found", http.StatusNotFound)
		return
	}

	// Send response with customer names
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customerNames)
}

// GetAllCustomerHandler retrieves all customers from the database
// @Summary Get all customers
// @Description This endpoint retrieves all customers from the database
// @Tags customers
// @Accept  json
// @Produce  json
// @Success 200 {array} Customer "List of all customers"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /customers [get]
func (app *Config) GetAllCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var customers []Customer

	// Fetch all customers from the database
	result := app.DB.Find(&customers)
	if result.Error != nil {
		http.Error(w, "Failed to retrieve customers", http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// OrderCustomersHandler retrieves all customers ordered by a specified field
// @Summary Get all customers ordered by a specific field
// @Description This endpoint retrieves customers from the database and orders them by a specified field.
//
//	If the "order_by" query parameter is not provided or is invalid, the customers will be ordered by "created_at" by default.
//
// @Tags customers
// @Accept  json
// @Produce  json
// @Param order_by query string false "Field to order by (customername, created_at, updated_at)"
// @Success 200 {array} Customer "List of ordered customers"
// @Failure 400 {object} ErrorResponse "Invalid order_by field"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /customers/order [get]
func (app *Config) OrderCustomersHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "order_by" query parameter
	orderBy := r.URL.Query().Get("order_by")

	// Allowed fields for ordering
	allowedFields := map[string]bool{
		"customername": true,
		"created_at":   true,
		"updated_at":   true,
	}

	// Default order by "created_at"
	if !allowedFields[orderBy] {
		orderBy = "created_at"
	}

	// Retrieve ordered customers from the database
	var customers []Customer
	result := app.DB.Order(orderBy + " DESC").Find(&customers)
	if result.Error != nil {
		http.Error(w, "Failed to fetch customers", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

// HealthCheckHandler checks the database connection using GORM
// @Summary Health check for the database connection
// @Description This endpoint checks if the database connection is successful by executing a lightweight query.
//
//	If the database connection is successful, it returns an "OK" response. Otherwise, it returns an error.
//
// @Tags health
// @Produce  plain
// @Success 200 {string} string "OK"
// @Failure 500 {string} string "Database connection failed"
// @Router /healthcheck [get]
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
// @Summary Create a new customer
// @Description This endpoint allows you to create a new customer with a hashed password. The customer must provide a unique customer name and mail address.
//
//	If the customer already exists (by customer name or mail address), the request will be rejected with a conflict error.
//
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body Customer true "Customer data"
// @Success 201 {object} map[string]string {"message": "Customer created successfully", "mailAddress": "string"}
// @Failure 400 {string} string "Invalid request body"
// @Failure 400 {string} string "Mail address cannot be empty"
// @Failure 409 {string} string "Customer already exists"
// @Failure 500 {string} string "Error hashing password"
// @Failure 500 {string} string "Error inserting customer"
// @Router /customers [post]
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
// @Summary Log in a customer
// @Description This endpoint allows customers to log in by verifying their credentials (customername and password). If credentials are valid, the login status will be updated to true.
//
//	If the customer does not exist or the password is incorrect, the request will be rejected with an unauthorized error.
//
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body Customer true "Customer login credentials"
// @Success 200 {object} map[string]string {"message": "Login successful", "loginStatus": "true"}
// @Failure 400 {string} string "Invalid request body"
// @Failure 401 {string} string "Customer not found"
// @Failure 401 {string} string "Invalid credentials"
// @Failure 500 {string} string "Database error"
// @Router /customers/login [post]
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
// @Summary Update customer password
// @Description This endpoint allows a customer to update their password. The customer must be identified by their customername. The new password is hashed before saving.
// @Tags customers
// @Accept  json
// @Produce  json
// @Param requestData body struct { Customername string `json:"customername"`; NewPassword string `json:"new_password"` } true "Customer password update data"
// @Success 200 {string} string "Password updated successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 404 {string} string "Customer not found"
// @Failure 500 {string} string "Database error"
// @Failure 500 {string} string "Error hashing password"
// @Router /customers/update-password [post]
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
// @Summary Get a customer by ID
// @Description This endpoint retrieves a customer record by the provided ID.
// @Tags customers
// @Accept  json
// @Produce  json
// @Param id query string true "Customer ID"
// @Success 200 {object} Customer "Customer retrieved successfully"
// @Failure 400 {string} string "Invalid request"
// @Failure 404 {string} string "Customer not found"
// @Failure 500 {string} string "Database error"
// @Router /customers [get]
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

// UpdateCustomerHandler updates customer details based on customername in the request body
// @Summary Update customer details
// @Description This endpoint updates customer information based on the customername provided in the request body.
// @Tags Customer
// @Accept  json
// @Produce  json
// @Param customer body struct {
//
//	    Customername string `json:"customername"`
//	    Password     string `json:"password,omitempty"`
//	    Note         string `json:"note,omitempty"`
//	    MailAddress  string `json:"mail_address,omitempty"`
//	} true "Customer details to update"
//
// @Success 200 {string} string "Customer updated successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 404 {string} string "Customer not found"
// @Failure 500 {string} string "Database error or internal error"
// @Router /customers/update [put]
func (app *Config) UpdateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Define a struct to hold the data from the request body
	var request struct {
		Customername string `json:"customername"`
		Password     string `json:"password,omitempty"`
		Note         string `json:"note,omitempty"`
		MailAddress  string `json:"mail_address,omitempty"`
	}

	// Decode the request body into the struct
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find the customer by customername
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

	// Update the customer's fields if provided in the request body
	if request.Password != "" {
		// Hash the new password
		hashedPassword, err := app.HashPassword(request.Password)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		customer.Password = hashedPassword
	}

	// Update the other fields (note, mail address) if provided
	if request.Note != "" {
		customer.Note = request.Note
	}
	if request.MailAddress != "" {
		customer.MailAddress = request.MailAddress
	}

	// Save the updated customer record
	if err := app.DB.Save(&customer).Error; err != nil {
		http.Error(w, "Failed to update customer", http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Customer updated successfully")
}

// DeactivateCustomerHandler deactivates a customer by customername
// @Summary Deactivate a customer by customername
// @Description This endpoint deactivates a customer's account by setting the "Activated" field to false.
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body struct {
//
//	    Customername string `json:"customername"`
//	} true "Customer name to deactivate"
//
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {string} string "Invalid request body or Customer name is required"
// @Failure 404 {string} string "Customer not found"
// @Failure 500 {string} string "Failed to deactivate customer"
// @Router /customers/deactivate [put]
func (app *Config) DeactivateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Define a struct to hold the data from the request body
	var request struct {
		Customername string `json:"customername"`
	}

	// Decode the request body into the struct
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure customer name is provided
	if request.Customername == "" {
		http.Error(w, "Customer name is required", http.StatusBadRequest)
		return
	}

	// Find the customer by customername
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
		"message": "Customer deactivated successfully",
	})
}

// ActivateCustomerHandler activates a customer by customername
// @Summary Activate a customer by customername
// @Description This endpoint activates a customer's account by setting the "Activated" field to true.
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body struct {
//
//	    Customername string `json:"customername"`
//	} true "Customer name to activate"
//
// @Success 200 {object} map[string]string "Success message"
// @Failure 400 {string} string "Invalid request body or Customer name is required"
// @Failure 404 {string} string "Customer not found"
// @Failure 500 {string} string "Failed to activate customer"
// @Router /customers/activate [put]
func (app *Config) ActivateCustomerHandler(w http.ResponseWriter, r *http.Request) {
	// Define a struct to hold the data from the request body
	var request struct {
		Customername string `json:"customername"`
	}

	// Decode the request body into the struct
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure customer name is provided
	if request.Customername == "" {
		http.Error(w, "Customer name is required", http.StatusBadRequest)
		return
	}

	// Find the customer by customername
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

	// Set Activated to true
	customer.Activated = true

	// Update the customer in the database
	if err := app.DB.Save(&customer).Error; err != nil {
		http.Error(w, "Failed to activate customer", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Customer activated successfully",
	})
}

// UpdateEmailHandler updates a customer's email address
// @Summary Update a customer's email address
// @Description This endpoint allows a customer to update their email address. The new email must be valid and not empty.
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body struct {
//
//	    Customername string `json:"customername"`
//	    NewEmail     string `json:"new_email"`
//	} true "Customer name and new email address"
//
// @Success 200 {string} string "Email updated successfully"
// @Failure 400 {string} string "Invalid request body or invalid email format"
// @Failure 404 {string} string "Customer not found"
// @Failure 500 {string} string "Failed to update email"
// @Router /customers/email [put]
func (app *Config) UpdateEmailHandler(w http.ResponseWriter, r *http.Request) {
	// Define a struct to hold the data from the request body
	var request struct {
		Customername string `json:"customername"`
		NewEmail     string `json:"new_email"`
	}

	// Decode the request body into the struct
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure customer name and email are provided
	if request.Customername == "" || request.NewEmail == "" {
		http.Error(w, "Customer name and new email are required", http.StatusBadRequest)
		return
	}

	// Check if the new email is valid (basic format validation)
	if !isValidEmail(request.NewEmail) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Find the customer by customername
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

	// Update the customer's email address
	customer.MailAddress = request.NewEmail
	if err := app.DB.Save(&customer).Error; err != nil {
		http.Error(w, "Failed to update email", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Email updated successfully",
	})
}

// UpdateNoteHandler updates the Note field for a customer
// @Summary Update a customer's note
// @Description This endpoint allows a user to update the note field of a customer by customername.
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customername body string true "Customer's name"
// @Param note body string true "New note for the customer"
// @Success 200 {string} string "Note updated successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 404 {string} string "Customer not found"
// @Failure 500 {string} string "Failed to update note"
// @Router /customers/note [put]
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

// InsertNoteHandler appends new text to the existing "Note" field for a customer
// @Summary Append a new note to an existing customer's note
// @Description This endpoint allows a user to append new text to the existing note of a customer by customername.
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customername body string true "Customer's name"
// @Param new_note body string true "New note to append"
// @Success 200 {string} string "Note appended successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 404 {string} string "Customer not found"
// @Failure 500 {string} string "Failed to update note"
// @Router /customers/note/append [put]
func (app *Config) InsertNoteHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Customername string `json:"customername"`
		NewNote      string `json:"new_note"`
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

	// Append the new note to the existing note (if any)
	if customer.Note != "" {
		customer.Note += "\n" // New line separator for readability
	}
	customer.Note += request.NewNote

	// Save the updated customer record with the new note
	if err := app.DB.Save(&customer).Error; err != nil {
		http.Error(w, "Failed to update note", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Note appended successfully",
	})
}

// DeleteCustomerHandler deletes a customer by their customername
// @Summary Delete a customer by their customername
// @Description This endpoint allows a user to delete a customer from the database by their customername.
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customername body string true "Customer name"
// @Success 200 {string} string "Customer deleted successfully"
// @Failure 400 {string} string "Customer name is required"
// @Failure 404 {string} string "Customer not found"
// @Failure 500 {string} string "Failed to delete customer"
// @Router /customers/delete [delete]
func (app *Config) DeleteCustomerHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Customername string `json:"customername"`
	}

	// Decode request body to get the customer name
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Ensure the customer name is provided
	if request.Customername == "" {
		http.Error(w, "Customer name is required", http.StatusBadRequest)
		return
	}

	// Find the customer by customername
	var customer Customer
	result := app.DB.Where("customername = ?", request.Customername).First(&customer)

	// If the customer is not found
	if result.RowsAffected == 0 {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	// Delete the customer
	result = app.DB.Delete(&customer)
	if result.Error != nil {
		http.Error(w, "Failed to delete customer", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Customer deleted successfully")
}

// Helper function to validate email format (simple validation)
func isValidEmail(email string) bool {

	// Basic validation for a common email format is used
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
