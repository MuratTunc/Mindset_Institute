package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// Constants for error and success messages
const (
	ErrInvalidRequestBody = "Invalid request body"
	ErrHashingPassword    = "Error hashing password"
	ErrInsertingUser      = "Error inserting sale"
	ErrUserNotFound       = "User not found"
	ErrInvalidCredentials = "Invalid credentials"
	UserCreatedSuccess    = "User created successfully"
	UserUpdatedSuccess    = "User updated successfully"
	UserDeletedSuccess    = "User deleted successfully"
	LoginSuccess          = "Login successful"
)

// HealthCheckHandler checks the database connection using GORM
// @Summary Health check
// @Description Check if the database connection is working
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 500 {string} string "Database connection failed"
// @Router /health [get]
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

// InsertSaleHandler creates a new sale record in the database
// @Summary Insert a new sale record
// @Description Creates a new sale record with the provided details
// @Tags Sales
// @Accept  json
// @Produce  json
// @Param salename body string true "Sale Name"
// @Param note body string false "Sale Note"
// @Success 200 {object} map[string]string {"message": "Sale record created successfully"}
// @Failure 400 {string} string "Invalid request body"
// @Failure 409 {string} string "Sale with this name already exists"
// @Failure 500 {string} string "Failed to create sale record"
// @Router /sale [post]
func (app *Config) InsertSaleHandler(w http.ResponseWriter, r *http.Request) {

	var request struct {
		Salename string `json:"salename"`
		Note     string `json:"note"`
	}

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if Salename already exists
	var existingSale Sale
	if err := app.DB.Where("salename = ?", request.Salename).First(&existingSale).Error; err == nil {
		http.Error(w, "Sale with this name already exists", http.StatusConflict)
		return
	} else if err != gorm.ErrRecordNotFound {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Create a new sale record
	newSale := Sale{
		Salename:  request.Salename,
		New:       true,
		Note:      request.Note,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save to DB
	if err := app.DB.Create(&newSale).Error; err != nil {
		http.Error(w, "Failed to create sale record", http.StatusInternalServerError)
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sale record created successfully",
	})
}

// DeleteSaleHandler deletes a sale record from the database
// @Summary Delete a sale record
// @Description Deletes a sale record identified by salename
// @Tags Sales
// @Accept  json
// @Produce  json
// @Param salename body string true "Sale Name"
// @Success 200 {object} map[string]string {"message": "Sale deleted successfully"}
// @Failure 400 {string} string "Invalid request body"
// @Failure 404 {string} string "Sale not found"
// @Failure 500 {string} string "Failed to delete sale record"
// @Router /sale [delete]
func (app *Config) DeleteSaleHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Salename string `json:"salename"`
	}

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the sale exists
	var sale Sale
	if err := app.DB.Where("salename = ?", request.Salename).First(&sale).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Sale not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Delete the sale record
	if err := app.DB.Delete(&sale).Error; err != nil {
		http.Error(w, "Failed to delete sale", http.StatusInternalServerError)
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Sale deleted successfully",
		"salename": request.Salename,
	})
}

// UpdateInCommunicationHandler updates the InCommunication field in the sale record
// @Summary Update the InCommunication field of a sale record
// @Description Updates the InCommunication status and appends a note to the sale record
// @Tags Sales
// @Accept  json
// @Produce  json
// @Param salename body string true "Sale Name"
// @Param in_communication body bool true "In Communication Status"
// @Param note body string false "Sale Note"
// @Success 200 {object} map[string]string {"message": "Sale record updated successfully"}
// @Failure 400 {string} string "Invalid request body"
// @Failure 404 {string} string "Sale not found"
// @Failure 500 {string} string "Failed to update sale record"
// @Router /sale/communication [put]
func (app *Config) UpdateInCommunicationHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Salename        string `json:"salename"`
		InCommunication bool   `json:"in_communication"`
		Note            string `json:"note"`
	}

	// Print incoming request data to logs for debugging
	fmt.Println("Received request to update sale record:")
	fmt.Printf("Salename: %s\n", request.Salename)
	fmt.Printf("InCommunication: %t\n", request.InCommunication)
	fmt.Printf("Note: %s\n", request.Note)

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find sale by salename
	var sale Sale
	if err := app.DB.Where("salename = ?", request.Salename).First(&sale).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Sale not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Update the fields based on request
	sale.InCommunication = request.InCommunication
	sale.New = !request.InCommunication // Set New to false if InCommunication is true

	// Append the new note to the existing note (on a new line)
	if request.Note != "" {
		sale.Note += "\n" + request.Note // Append the new note on a new line
	}

	// Update the sale record and log it
	if err := app.DB.Save(&sale).Error; err != nil {
		http.Error(w, "Failed to update sale record", http.StatusInternalServerError)
		return
	}

	// Print the updated sale data for debugging
	fmt.Println("Updated sale record:")
	fmt.Printf("Salename: %s\n", sale.Salename)
	fmt.Printf("InCommunication: %t\n", sale.InCommunication)
	fmt.Printf("Note: %s\n", sale.Note)

	// Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sale record updated successfully",
	})
}

// UpdateDealHandler updates the Deal field in the sale record and modifies other fields as required
// @Summary Update a sale record's deal status and associated information
// @Description Updates the Deal field of an existing sale record by salename and optionally appends a note
// @Tags Sale
// @Accept json
// @Produce json
// @Param salename body string true "Salename"
// @Param deal body bool true "Deal Status (true for deal, false for not)"
// @Param note body string false "Note to append to the sale record"
// @Success 200 {object} map[string]string {"message": "Sale record updated successfully"}
// @Failure 400 {object} map[string]string {"error": "Invalid request body"}
// @Failure 404 {object} map[string]string {"error": "Sale not found"}
// @Failure 500 {object} map[string]string {"error": "Database error"}
// @Router /sales/update-deal [put]
func (app *Config) UpdateDealHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Salename string `json:"salename"`
		Deal     bool   `json:"deal"`
		Note     string `json:"note"`
	}

	// Print incoming request data to logs for debugging
	fmt.Println("Received request to update sale record:")
	fmt.Printf("Salename: %s\n", request.Salename)
	fmt.Printf("Deal: %t\n", request.Deal)
	fmt.Printf("Note: %s\n", request.Note)

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find sale by salename
	var sale Sale
	if err := app.DB.Where("salename = ?", request.Salename).First(&sale).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Sale not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Update the fields based on request
	sale.Deal = request.Deal
	if request.Deal {
		sale.New = false
		sale.InCommunication = false
	}

	// Append the new note to the existing note (on a new line)
	if request.Note != "" {
		sale.Note += "\n" + request.Note // Append the new note on a new line
	}

	// Update the UpdatedAt field to the current time
	sale.UpdatedAt = time.Now()

	// Update the sale record and log it
	if err := app.DB.Save(&sale).Error; err != nil {
		http.Error(w, "Failed to update sale record", http.StatusInternalServerError)
		return
	}

	// Print the updated sale data for debugging
	fmt.Println("Updated sale record:")
	fmt.Printf("Salename: %s\n", sale.Salename)
	fmt.Printf("Deal: %t\n", sale.Deal)
	fmt.Printf("New: %t\n", sale.New)
	fmt.Printf("InCommunication: %t\n", sale.InCommunication)
	fmt.Printf("Note: %s\n", sale.Note)
	fmt.Printf("UpdatedAt: %s\n", sale.UpdatedAt)

	// Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sale record updated successfully",
	})
}

// UpdateClosedHandler updates the "Closed" field in the sale record and modifies other fields as required
// @Summary Close a sale record and update associated information
// @Description Updates the "Closed" field of an existing sale record by salename and optionally appends a note
// @Tags Sale
// @Accept json
// @Produce json
// @Param salename body string true "Salename"
// @Param note body string false "Note to append to the sale record"
// @Success 200 {object} map[string]string {"message": "Sale record closed successfully"}
// @Failure 400 {object} map[string]string {"error": "Invalid request body"}
// @Failure 404 {object} map[string]string {"error": "Sale not found"}
// @Failure 500 {object} map[string]string {"error": "Database error"}
// @Router /sales/update-closed [put]
func (app *Config) UpdateClosedHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Salename string `json:"salename"`
		Note     string `json:"note"`
	}

	// Print incoming request data to logs for debugging
	fmt.Println("Received request to update sale record:")
	fmt.Printf("Salename: %s\n", request.Salename)
	fmt.Printf("Note: %s\n", request.Note)

	// Decode request body
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find sale by salename
	var sale Sale
	if err := app.DB.Where("salename = ?", request.Salename).First(&sale).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Sale not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Update the fields based on request
	sale.Closed = true
	sale.New = false
	sale.InCommunication = false
	sale.Deal = false

	// Append the note text, if provided
	if request.Note != "" {
		sale.Note += "\n" + request.Note
	}

	// Update the UpdatedAt field
	sale.UpdatedAt = time.Now()

	// Update the sale record in the database
	if err := app.DB.Save(&sale).Error; err != nil {
		http.Error(w, "Failed to update sale record", http.StatusInternalServerError)
		return
	}

	// Print the updated sale data for debugging
	fmt.Println("Updated sale record:")
	fmt.Printf("Salename: %s\n", sale.Salename)
	fmt.Printf("Closed: %t\n", sale.Closed)
	fmt.Printf("Note: %s\n", sale.Note)

	// Success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Sale record closed successfully",
	})
}
