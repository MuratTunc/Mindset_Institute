package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

// connectToDB function to connect to PostgreSQL using constants
func connectToDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DBHost, DBUser, DBPassword, DBName, DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		fmt.Println("DATABASE connection success!")
	}

	// AutoMigrate to create tables
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db, nil
}

func main() {

	PrintEnvVariables()

	db, err := connectToDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	var app = Config{DB: db}

	ServicePort := os.Getenv("AUTHENTICATION_SERVICE_PORT")
	ServiceName := os.Getenv("AUTHENTICATION_SERVICE_NAME")

	if ServicePort == "" || ServiceName == "" {
		log.Fatal("Error: Authentication Service environment variables are not set")
	}

	fmt.Printf("%s is running on port: %s", ServiceName, ServicePort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", ServicePort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic()
	}
}
