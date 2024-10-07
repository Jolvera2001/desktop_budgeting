package models

type Category struct {
	Base
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID uint   `json:"userId"`
	Name   string `json:"category"`
}
