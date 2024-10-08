package models

import "gorm.io/gorm"

type Budget struct {
	gorm.Model
	User       User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID     uint     `json:"user_id"`
	CategoryID uint     `json:"category"`
	Name       string   `json:"name"`
	Amount     float64  `json:"amount"`
}
