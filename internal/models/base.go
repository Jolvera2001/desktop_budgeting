package models

import "time"

type Base struct {
	ID        uint      `gorm:"primaryKey" json:"_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
