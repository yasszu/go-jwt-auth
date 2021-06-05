package entity

import (
	"time"
)

type Account struct {
	ID           uint `gorm:"primaryKey"`
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
