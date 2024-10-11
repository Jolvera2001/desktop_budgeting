package models

import (
	"time"
)

type Transaction struct {
	BaseModel
	Budget      Budget    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BudgetID    uint      `json:"batch_id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date" gorm:"index"`
}
