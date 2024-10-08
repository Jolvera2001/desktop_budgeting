package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email        string        `json:"email"`
	Name         string        `json:"name"`
	Transactions []Transaction `gorm:"foreignKey:UserID"`
	Categories   []Category    `gorm:"foreignKey:UserID"`
	Budgets      []Budget      `gorm:"foreignKey:UserID"`
}
