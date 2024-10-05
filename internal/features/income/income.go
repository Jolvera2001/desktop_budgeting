package income

import (
	"time"
)

type Income struct {
	ID         int64     `json:"_id"`
	UserID     int64     `json:"user_id"`
	CategoryID int64     `json:"category"`
	Amount     float64   `json:"amount"`
	IsRegular  bool      `json:"is_regular"`
	Date       time.Time `json:"date"`
}
