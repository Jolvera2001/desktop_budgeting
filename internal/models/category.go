package models

type Category struct {
	ID     uint   `json:"_id" gorm:"primaryKey"`
	UserID uint   `json:"userId"`
	Name   string `json:"category"`
}
