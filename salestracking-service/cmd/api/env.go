package main

import (
	"fmt"
	"log"
	"os"
)

var (
	DBHost      = os.Getenv("SALESTRACKING_POSTGRES_DB_HOST")
	DBUser      = os.Getenv("SALESTRACKING_POSTGRES_DB_USER")
	DBPassword  = os.Getenv("SALESTRACKING_POSTGRES_DB_PASSWORD")
	DBName      = os.Getenv("SALESTRACKING_POSTGRES_DB_NAME")
	DBPort      = os.Getenv("SALESTRACKING_POSTGRES_DB_PORT")
	ServicePort = os.Getenv("SALESTRACKING_SERVICE_PORT")
	ServiceName = os.Getenv("SALESTRACKING_SERVICE_NAME")
)

// PrintEnvVariables prints all environment variables for debugging
func PrintEnvVariables() {
	fmt.Println("🔧 Loaded Environment Variables-SALESTRACKING_SERVICE")
	fmt.Printf("DBUser: %s\n", DBUser)
	fmt.Printf("DBPassword: %s\n", DBPassword)
	fmt.Printf("DBName: %s\n", DBName)
	fmt.Printf("DBPort: %s\n", DBPort)
	fmt.Printf("DBHost: %s\n", DBHost)
	fmt.Printf("ServicePort: %s\n", ServicePort)
	fmt.Printf("ServiceName: %s\n", ServiceName)

	// Ensure SERVICE environment variables are set
	if ServicePort == "" || ServiceName == "" {
		log.Fatal("❌ Error: Missing environment variables for SALESTRACKING_SERVICE")
	}
}
