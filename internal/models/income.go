package models

import (
	"time"
)

type IncomeType string

const (
	Regular   IncomeType = "regular"
	OneTime   IncomeType = "one_time"
	Recurring IncomeType = "recurring"
)

type Income struct {
	BaseModel
	User       User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID     uint       `json:"user_id"`
	Amount     float64    `json:"amount"`
	IncomeType IncomeType `gorm:"type:varchar(20)" json:"income_type"`
	Date       time.Time  `json:"date" gorm:"index"`
}
