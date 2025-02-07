package main

import "time"

// Customer model for GORM
type Customer struct {
	ID           uint      `gorm:"primaryKey"`
	Customername string    `gorm:"unique;not null"`
	MailAddress  string    `gorm:"unique;not null"`
	Password     string    `gorm:"not null"`
	Activated    bool      `gorm:"default:false"`
	LoginStatus  bool      `gorm:"default:false"`
	Note         string    `gorm:"type:text"` // Field to store text information
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
