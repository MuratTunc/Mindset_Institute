package main

import (
	"fmt"
	"os"
)

var (
	DBHost      = os.Getenv("DB_HOST")     // "auth-db"
	DBUser      = os.Getenv("DB_USER")     // "auth_user"
	DBPassword  = os.Getenv("DB_PASSWORD") // "auth_password"
	DBName      = os.Getenv("DB_NAME")     // "auth_db"
	DBPort      = os.Getenv("DB_PORT")     // "5432"
	ServicePort = os.Getenv("AUTHENTICATION_SERVICE_PORT")
	ServiceName = os.Getenv("AUTHENTICATION_SERVICE_NAME")
)

// PrintEnvVariables prints all environment variables for debugging
func PrintEnvVariables() {
	fmt.Println("🔧 Loaded Environment Variables:")
	fmt.Printf("DBUser: %s\n", DBUser)
	fmt.Printf("DBPassword: %s\n", DBPassword)
	fmt.Printf("DBName: %s\n", DBName)
	fmt.Printf("DBPort: %s\n", DBPort)
	fmt.Printf("DBHost: %s\n", DBHost)
	fmt.Printf("ServicePort: %s\n", ServicePort)
	fmt.Printf("ServiceName: %s\n", ServiceName)
}
