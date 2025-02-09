package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Setup GORM mock database using sqlmock (GORM v2)
func setupMockGORMDB() (*gorm.DB, sqlmock.Sqlmock) {
	// Create the mock DB using sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("failed to create mock db")
	}

	// Convert the *sql.DB into a *gorm.DB (GORM v2)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		panic("failed to create GORM DB")
	}

	return gormDB, mock
}

func TestHealthCheckHandler(t *testing.T) {
	// Initialize the app with a mock database using the setupMockGORMDB function
	db, mock := setupMockGORMDB()
	app := &Config{
		DB: db,
	}

	// Create a new HTTP request for health check
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	// Create a new ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the HealthCheckHandler with the request and ResponseRecorder
	handler := http.HandlerFunc(app.HealthCheckHandler)
	handler.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", rr.Code)
	}

	// Check if the response body contains "OK"
	expected := "OK"
	if rr.Body.String() != expected {
		t.Errorf("Expected response body to be %v, got %v", expected, rr.Body.String())
	}

	// Ensure all mock expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unmet expectations: %s", err)
	}
}
