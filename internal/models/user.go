package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	Incomes []Income `gorm:"foreignKey:UserID" json:"incomes"`
	Budgets []Budget `gorm:"foreignKey:UserID" json:"budgets"`
}
