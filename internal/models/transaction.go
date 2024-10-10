package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model  `json:",inline"`
	Budget      Budget    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BudgetID    uint      `json:"batch_id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date" gorm:"index"`
}
