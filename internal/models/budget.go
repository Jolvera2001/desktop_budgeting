package models

type Budget struct {
	Base
	UserID     uint    `json:"user_id"`
	CategoryID uint    `json:"category"`
	Name       string  `json:"name"`
	Amount     float64 `json:"amount"`
}
