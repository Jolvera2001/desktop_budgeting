package models

type Budget struct {
	ID         uint  `json:"_id" gorm:"primaryKey"`
	UserID     uint  `json:"user_id"`
	CategoryID uint   `json:"category"`
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
}
