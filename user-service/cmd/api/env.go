package main

import (
	"fmt"
	"log"
	"os"
)

var (
	DBHost      = os.Getenv("USER_POSTGRES_DB_HOST")
	DBUser      = os.Getenv("USER_POSTGRES_DB_USER")
	DBPassword  = os.Getenv("USER_POSTGRES_DB_PASSWORD")
	DBName      = os.Getenv("USER_POSTGRES_DB_NAME")
	DBPort      = os.Getenv("USER_POSTGRES_DB_PORT")
	ServicePort = os.Getenv("USER_SERVICE_PORT")
	ServiceName = os.Getenv("USER_SERVICE_NAME")
)

// PrintEnvVariables prints all environment variables for debugging
func PrintEnvVariables() {
	fmt.Println("🔧 Loaded Environment Variables:")
	fmt.Printf("USER_POSTGRES_DBUser: %s\n", DBUser)
	fmt.Printf("USER_POSTGRES_DBPassword: %s\n", DBPassword)
	fmt.Printf("USER_POSTGRES_DBName: %s\n", DBName)
	fmt.Printf("USER_POSTGRES_DBPort: %s\n", DBPort)
	fmt.Printf("USER_POSTGRES_DBHost: %s\n", DBHost)
	fmt.Printf("USER_POSTGRES_ServicePort: %s\n", ServicePort)
	fmt.Printf("USER_POSTGRES_ServiceName: %s\n", ServiceName)

	// Ensure SERVICE environment variables are set
	if ServicePort == "" || ServiceName == "" {
		log.Fatal("❌ Error: Missing environment variables for service")
	}
}
