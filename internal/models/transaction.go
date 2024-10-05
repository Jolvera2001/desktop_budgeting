package models

import "time"

type Transaction struct {
	ID          uint      `json:"_id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	BudgetID    uint      `json:"batch_id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	DateCreated time.Time `json:"date"`
}
