package models

import "time"

type Transaction struct {
	Base
	UserID      uint      `json:"user_id"`
	BudgetID    uint      `json:"batch_id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date"`
}
