package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// Test to check that the environment variables are loaded correctly
func TestLoadEnvVariables(t *testing.T) {

	err := godotenv.Load("../../../build-tools/.env")
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	DBHost = os.Getenv("USER_POSTGRES_DB_HOST")
	DBUser = os.Getenv("USER_POSTGRES_DB_USER")
	DBPassword = os.Getenv("USER_POSTGRES_DB_PASSWORD")
	DBName = os.Getenv("USER_POSTGRES_DB_NAME")
	ServicePort = os.Getenv("USER_SERVICE_PORT")
	ServiceName = os.Getenv("USER_SERVICE_NAME")

	fmt.Println("🔧 Loaded Environment Variables-USER_SERVICE")
	fmt.Printf("DBUser: %s\n", DBUser)
	fmt.Printf("DBPassword: %s\n", DBPassword)
	fmt.Printf("DBName: %s\n", DBName)
	fmt.Printf("DBHost: %s\n", DBHost)
	fmt.Printf("ServicePort: %s\n", ServicePort)
	fmt.Printf("ServiceName: %s\n", ServiceName)

	// Verify the environment variables are set correctly
	assert.Equal(t, "user-db", DBHost, "DBHost should be localhost")
	assert.Equal(t, "user", DBUser, "DBUser should be user")
	assert.Equal(t, "user_password", DBPassword, "DBPassword should be password")
	assert.Equal(t, "user_db", DBName, "DBName should be userdb")
	assert.Equal(t, "8080", ServicePort, "ServicePort should be 8080")
	assert.Equal(t, "USER-SERVICE", ServiceName, "ServiceName should be user-service")
}
