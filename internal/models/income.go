package models

import (
	"time"
)

type Income struct {
	ID         uint     `json:"_id" gorm:"primaryKey"`
	UserID     uint     `json:"user_id"`
	CategoryID uint     `json:"category"`
	Amount     float64   `json:"amount"`
	IsRegular  bool      `json:"is_regular"`
	Date       time.Time `json:"date"`
}
