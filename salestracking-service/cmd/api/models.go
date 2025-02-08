package main

import "time"

// Sale model for GORM
type Sale struct {
	ID              uint      `gorm:"primaryKey"`
	Salename        string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	New             bool      `gorm:"default:false"`
	InCommunication bool      `gorm:"default:false"`
	Deal            bool      `gorm:"default:false"`
	Closed          bool      `gorm:"default:false"`
	Note            string    `gorm:"type:text"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
