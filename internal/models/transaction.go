package models

import "time"

type Transaction struct {
	Base
	User        User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Budget      Budget    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID      uint      `json:"user_id"`
	BudgetID    uint      `json:"batch_id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	Date        time.Time `json:"date" gorm:"index"`
}
