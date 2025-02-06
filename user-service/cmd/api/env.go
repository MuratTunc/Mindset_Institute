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
	JWTSecret   = os.Getenv("USER_SERVICE_JWT_SECRET")
)

// PrintEnvVariables prints all environment variables for debugging
func PrintEnvVariables() {
	fmt.Println("🔧 Loaded Environment Variables-USER_SERVICE")
	fmt.Printf("DBUser: %s\n", DBUser)
	fmt.Printf("DBPassword: %s\n", DBPassword)
	fmt.Printf("DBName: %s\n", DBName)
	fmt.Printf("DBPort: %s\n", DBPort)
	fmt.Printf("DBHost: %s\n", DBHost)
	fmt.Printf("ServicePort: %s\n", ServicePort)
	fmt.Printf("ServiceName: %s\n", ServiceName)
	fmt.Printf("JWTSecret: %s\n", JWTSecret)

	// Ensure SERVICE environment variables are set
	if ServicePort == "" || ServiceName == "" {
		log.Fatal("❌ Error: Missing environment variables for USER_SERVICE")
	}
}
