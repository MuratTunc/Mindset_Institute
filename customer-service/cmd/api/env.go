package main

import (
	"fmt"
	"log"
	"os"
)

var (
	DBHost      = os.Getenv("CUSTOMER_POSTGRES_DB_HOST")
	DBCustomer  = os.Getenv("CUSTOMER_POSTGRES_DB_CUSTOMER")
	DBPassword  = os.Getenv("CUSTOMER_POSTGRES_DB_PASSWORD")
	DBName      = os.Getenv("CUSTOMER_POSTGRES_DB_NAME")
	DBPort      = os.Getenv("CUSTOMER_POSTGRES_DB_PORT")
	ServicePort = os.Getenv("CUSTOMER_SERVICE_PORT")
	ServiceName = os.Getenv("CUSTOMER_SERVICE_NAME")
)

// PrintEnvVariables prints all environment variables for debugging
func PrintEnvVariables() {
	fmt.Println("🔧 Loaded Environment Variables-CUSTOMER_SERVICE")
	fmt.Printf("DBCustomer: %s\n", DBCustomer)
	fmt.Printf("DBPassword: %s\n", DBPassword)
	fmt.Printf("DBName: %s\n", DBName)
	fmt.Printf("DBPort: %s\n", DBPort)
	fmt.Printf("DBHost: %s\n", DBHost)
	fmt.Printf("ServicePort: %s\n", ServicePort)
	fmt.Printf("ServiceName: %s\n", ServiceName)

	// Ensure SERVICE environment variables are set
	if ServicePort == "" || ServiceName == "" {
		log.Fatal("❌ Error: Missing environment variables for CUSTOMER_SERVICE")
	}
}
