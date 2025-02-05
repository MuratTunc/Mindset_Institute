package main

import "time"

// User model for GORM
type User struct {
	ID          uint      `gorm:"primaryKey"`
	Username    string    `gorm:"unique;not null"`
	MailAddress string    `gorm:"unique;not null"`
	Password    string    `gorm:"not null"`
	Activated   bool      `gorm:"default:false"`
	LoginStatus bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
