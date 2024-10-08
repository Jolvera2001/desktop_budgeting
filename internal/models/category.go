package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID uint   `json:"userId"`
	Name   string `json:"category"`
}
