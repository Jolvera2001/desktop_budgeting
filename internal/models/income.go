package models

import (
	"time"

	"gorm.io/gorm"
)

type Income struct {
	gorm.Model
	User       User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category   Category  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID     uint      `json:"user_id"`
	CategoryID uint      `json:"category"`
	Amount     float64   `json:"amount"`
	IsRegular  bool      `json:"is_regular"`
	Date       time.Time `json:"date" gorm:"index"`
}
