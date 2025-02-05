package main

import (
	"fmt"
	"os"
)

var (
	DBHost      = os.Getenv("CUSTOMER_POSTGRES_DB_HOST")
	DBUser      = os.Getenv("CUSTOMER_POSTGRES_DB_USER")
	DBPassword  = os.Getenv("CUSTOMER_POSTGRES_DB_PASSWORD")
	DBName      = os.Getenv("CUSTOMER_POSTGRES_DB_NAME")
	DBPort      = os.Getenv("CUSTOMER_POSTGRES_DB_PORT")
	ServicePort = os.Getenv("CUSTOMER_SERVICE_PORT")
	ServiceName = os.Getenv("CUSTOMER_SERVICE_NAME")
)

// PrintEnvVariables prints all environment variables for debugging
func PrintEnvVariables() {
	fmt.Println("🔧 Loaded Environment Variables:")
	fmt.Printf("CUSTOMER_POSTGRES_DBUser: %s\n", DBUser)
	fmt.Printf("CUSTOMER_POSTGRES_DBPassword: %s\n", DBPassword)
	fmt.Printf("CUSTOMER_POSTGRES_DBName: %s\n", DBName)
	fmt.Printf("CUSTOMER_POSTGRES_DBPort: %s\n", DBPort)
	fmt.Printf("CUSTOMER_POSTGRES_DBHost: %s\n", DBHost)
	fmt.Printf("CUSTOMER_POSTGRES_ServicePort: %s\n", ServicePort)
	fmt.Printf("CUSTOMER_POSTGRES_ServiceName: %s\n", ServiceName)
}
