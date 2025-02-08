package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

// connectToDB retries connecting to PostgreSQL until it succeeds or fails after retries
func connectToDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DBHost, DBUser, DBPassword, DBName, DBPort)

	var db *gorm.DB
	var err error

	// Retry logic: Try connecting 10 times with a 5-second delay
	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("✅ DATABASE connection success!")
			break
		}
		fmt.Printf("⏳ Attempt %d: Waiting for database to be ready...\n", i)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("❌ Failed to connect to database after retries: %v", err)
	}

	// AutoMigrate to create tables
	err = db.AutoMigrate(&Sale{})
	if err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}

	return db, nil
}

func main() {
	// Print environment variables for debugging
	PrintEnvVariables()

	db, err := connectToDB()
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	fmt.Printf("🚀 %s is running on port: %s\n", ServiceName, ServicePort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", ServicePort),
		Handler: (&Config{DB: db}).routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("💥 Server failed to start: %v", err)
	}
}
