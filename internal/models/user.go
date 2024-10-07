package models

type User struct {
	Base
	Email        string        `json:"email"`
	Name         string        `json:"name"`
	Transactions []Transaction `gorm:"foreignKey:UserID"`
	Categories   []Category    `gorm:"foreignKey:UserID"`
}
