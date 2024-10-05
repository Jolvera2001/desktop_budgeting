package models

import "time"

type Transaction struct {
	ID          int       `json:"_id"`
	UserID      int       `json:"user_id"`
	BudgetID    int       `json:"batch_id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	DateCreated time.Time `json:"date"`
}
