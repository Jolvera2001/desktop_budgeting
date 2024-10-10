package models

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:",inline"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Password   string   `json:"password"`
	Incomes    []Income `gorm:"foreignKey:UserID" json:"incomes"`
	Budgets    []Budget `gorm:"foreignKey:UserID" json:"budgets"`
}

type UserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
