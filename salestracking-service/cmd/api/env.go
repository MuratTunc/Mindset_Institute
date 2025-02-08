package main

import (
	"fmt"
	"log"
	"os"
)

// Define service name
const ServiceNamePrefix = "SALESTRACKING"

// Load environment variables dynamically using ServiceName
var (
	DBHost      = os.Getenv(ServiceNamePrefix + "_POSTGRES_DB_HOST")
	DBUser      = os.Getenv(ServiceNamePrefix + "_POSTGRES_DB_USER")
	DBPassword  = os.Getenv(ServiceNamePrefix + "_POSTGRES_DB_PASSWORD")
	DBName      = os.Getenv(ServiceNamePrefix + "_POSTGRES_DB_NAME")
	ServicePort = os.Getenv(ServiceNamePrefix + "_SERVICE_PORT")
	ServiceName = os.Getenv(ServiceNamePrefix + "_SERVICE_NAME")
)

// Set DBPort explicitly to 5432 inside the container
const DBPort = "5432"

// PrintEnvVariables prints all environment variables for debugging
func PrintEnvVariables() {
	fmt.Println("🔧 Loaded Environment Variables -" + ServiceNamePrefix + "_SERVICE")
	fmt.Printf("DBHost: %s\n", DBHost)
	fmt.Printf("DBUser: %s\n", DBUser)
	fmt.Printf("DBPassword: %s\n", DBPassword)
	fmt.Printf("DBName: %s\n", DBName)
	fmt.Printf("DBPort: %s\n", DBPort)
	fmt.Printf("ServicePort: %s\n", ServicePort)
	fmt.Printf("ServiceName: %s\n", ServiceName)

	// Ensure all required environment variables are set
	missingEnvVars := false
	if DBHost == "" || DBUser == "" || DBPassword == "" || DBName == "" {
		fmt.Println("❌ Error: Missing required database environment variables")
		missingEnvVars = true
	}
	if ServicePort == "" || ServiceName == "" {
		fmt.Println("❌ Error: Missing required service environment variables")
		missingEnvVars = true
	}

	if missingEnvVars {
		log.Fatal("❌ Exiting due to missing environment variables.")
	}
}
