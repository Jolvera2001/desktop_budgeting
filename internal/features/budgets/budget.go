package budgets

import "time"

type Budget struct {
	ID       int       `json:"_id"`
	UserID   int       `json:"user_id"`
	Category string    `json:"category"`
	Amount   float64   `json:"amount"`
	Date     time.Time `json:"date"`
}
