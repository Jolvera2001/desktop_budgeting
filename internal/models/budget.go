package models

type Budget struct {
	BaseModel
	User         User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID       uint          `json:"user_id"`
	Name         string        `json:"name"`
	Amount       float64       `json:"amount"`
	Transactions []Transaction `json:"transactions"`
}

type BudgetDto struct {
	UserID       uint          `json:"user_id"`
	Name         string        `json:"name"`
	Amount       float64       `json:"amount"`
}
