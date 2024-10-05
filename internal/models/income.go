package models

import (
	"time"
)

type Income struct {
	Base
	UserID     uint      `json:"user_id"`
	CategoryID uint      `json:"category"`
	Amount     float64   `json:"amount"`
	IsRegular  bool      `json:"is_regular"`
	Date       time.Time `json:"date"`
}
